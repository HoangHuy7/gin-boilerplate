package service

import (
	"context"
	"encoding/json"
	"fmt"
	"monorepo/apps/gas/app/config"
	"monorepo/apps/gas/app/database"
	"monorepo/shares/entities/mekyra_db"
	"monorepo/shares/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type LogService struct {
	db     *database.DataSources
	logger *zap.Logger
}

func NewLogService(db *database.DataSources, lg *zap.Logger) *LogService {
	return &LogService{
		db:     db,
		logger: lg,
	}
}

// LogActionInput contains all parameters for logging an action
type LogActionInput struct {
	Action    string      // CREATE, UPDATE, DELETE, etc.
	Status    string      // SUCCESS, FAILED, etc.
	Feature   string      // PRODUCT, CUSTOMER, ORDER, etc.
	OldData   interface{} // Data before the action (optional)
	NewData   interface{} // Data after the action (optional)
	CreatedBy string      // User who performed the action
	IpAddress string      // Client IP address
	UserAgent string      // Client user agent
	RequestId string      // Request ID for tracing
}

// LogAction creates a log entry in the database
func (s *LogService) LogAction(ctx context.Context, input LogActionInput) error {
	oldDataBytes, err := json.Marshal(input.OldData)
	if err != nil {
		s.logger.Error("Failed to marshal old data", zap.Error(err))
		oldDataBytes = []byte("{}")
	}

	newDataBytes, err := json.Marshal(input.NewData)
	if err != nil {
		s.logger.Error("Failed to marshal new data", zap.Error(err))
		newDataBytes = []byte("{}")
	}

	logEntry := &mekyra_db.Mkrtb_Logs{
		Action:    input.Action,
		Status:    input.Status,
		Feature:   input.Feature,
		OldData:   oldDataBytes,
		NewData:   newDataBytes,
		CreatedBy: input.CreatedBy,
		IpAddress: input.IpAddress,
		UserAgent: input.UserAgent,
		RequestId: input.RequestId,
	}

	org, getOrgErr := utils.GetOrg(ctx)
	if getOrgErr {
		s.logger.Error("Failed to get organization from context")
		return fmt.Errorf("failed to get organization from context")
	}

	tenancy, tenancyErr := config.GetTenancy(org)
	if tenancyErr {
		s.logger.Error("Failed to get tenancy")
		return fmt.Errorf("failed to get tenancy")
	}

	err = database.WithTenant(
		s.db.Mekyra_db,
		ctx,
		tenancy,
		func(tx *gorm.DB) error {
			return tx.Create(logEntry).Error
		},
	)

	if err != nil {
		s.logger.Error("Failed to create log entry", zap.Error(err))
		return fmt.Errorf("failed to create log entry: %w", err)
	}

	return nil
}

// LogSuccess is a helper method to log successful actions
func (s *LogService) LogSuccess(ctx context.Context, action, feature string, oldData, newData interface{}) {
	input := LogActionInput{
		Action:    action,
		Status:    "SUCCESS",
		Feature:   feature,
		OldData:   oldData,
		NewData:   newData,
		CreatedBy: "system",
	}
	_ = s.LogAction(ctx, input)
}

// LogError is a helper method to log failed actions
func (s *LogService) LogError(ctx context.Context, action, feature string, oldData interface{}, errorMsg string) {
	input := LogActionInput{
		Action:    action,
		Status:    "FAILED",
		Feature:   feature,
		OldData:   oldData,
		NewData:   map[string]string{"error": errorMsg},
		CreatedBy: "system",
	}
	_ = s.LogAction(ctx, input)
}

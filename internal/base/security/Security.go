package security

import (
	"context"
	"fmt"
	"monorepo/apps/gas/service"
	"monorepo/internal/dto"
	"monorepo/internal/logger"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type Security struct {
	Logger   *zap.Logger
	tenant   *dto.OIDC
	Oauth    oauth2.Config
	Verifier *oidc.IDTokenVerifier
}

func NewSecurity(tenant *dto.OIDC, goLogger *logger.GoLogger, r *service.RedisService) *Security {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, tenant.Issuer)
	if err != nil {
		panic(err)
	}

	verifier := provider.Verifier(&oidc.Config{
		SkipIssuerCheck:   true,
		SkipClientIDCheck: true,
	})

	oauth2Config := oauth2.Config{
		ClientID:     tenant.ClientID,
		ClientSecret: tenant.ClientSecret,
		RedirectURL:  "",
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "roles"},
	}

	return &Security{
		Logger:   goLogger.Zap,
		tenant:   tenant,
		Oauth:    oauth2Config,
		Verifier: verifier,
	}
}

func (this Security) BeforeFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractToken(c)
		if err != nil {
			this.Logger.Error("Error extracting token", zap.Error(err))
			c.AbortWithStatus(401)
			return
		}

		idToken, err := this.Verifier.Verify(c.Request.Context(), token)

		if err != nil {
			this.Logger.Error("Error verifying token", zap.Error(err))
			c.AbortWithStatus(401)
			return
		}
		this.Logger.Info("Token verified", zap.Any("token", idToken))
		c.Next()
	}
}

func extractToken(c *gin.Context) (string, error) {
	h := c.GetHeader("Authorization")
	if h == "" {
		return "", fmt.Errorf("missing token")
	}
	return strings.TrimPrefix(h, "Bearer "), nil
}

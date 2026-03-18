package security

import (
	"fmt"
	"monorepo/apps/gas/service"
	"monorepo/internal/dto"
	"monorepo/internal/logger"
	"monorepo/shares/utils"
	"strings"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

// CasdoorOrg holds SDK instance for each organization
type CasdoorOrg struct {
	Name string
	SDK  *casdoorsdk.Client
}

type Security struct {
	Logger        *zap.Logger
	Organizations map[string]*CasdoorOrg // key: org name (e.g., "directusO")
}

func NewSecurity(config *dto.CasdoorConfig, goLogger *logger.GoLogger, r *service.RedisService) *Security {
	orgs := make(map[string]*CasdoorOrg)

	for name, orgConfig := range config.Organizations {
		client := casdoorsdk.NewClient(
			orgConfig.ServerURL,
			orgConfig.ClientID,
			orgConfig.ClientSecret,
			orgConfig.Certificate,
			orgConfig.Organization,
			orgConfig.Application,
		)

		orgs[name] = &CasdoorOrg{
			Name: name,
			SDK:  client,
		}
	}

	return &Security{
		Logger:        goLogger.Zap,
		Organizations: orgs,
	}
}

func (s *Security) BeforeFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		//if true {
		//	c.Next()
		//	return
		//}
		tokenString, err := extractToken(c)
		if err != nil {
			s.Logger.Error("Missing token", zap.Error(err))
			c.AbortWithStatusJSON(401, gin.H{"error": "missing token"})
			return
		}

		// Parse JWT without verification to get the organization (owner)
		// Using jwt.Parse with nil keyfunc to skip verification
		token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &casdoorsdk.Claims{})
		if err != nil {
			s.Logger.Error("Failed to parse token", zap.Error(err))
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token format"})
			return
		}

		// Extract claims
		claims, ok := token.Claims.(*casdoorsdk.Claims)
		if !ok {
			s.Logger.Error("Invalid token claims")
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token claims"})
			return
		}

		// Get organization from JWT owner field
		orgName := claims.Owner
		if orgName == "" {
			s.Logger.Error("Token missing owner field")
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token: missing owner"})
			return
		}

		// Get org config
		org, exists := s.Organizations[orgName]
		if !exists {
			s.Logger.Error("Unknown organization", zap.String("org", orgName))
			c.AbortWithStatusJSON(401, gin.H{"error": "unknown organization"})
			return
		}

		// Verify JWT with the correct organization's certificate
		verifiedClaims, err := org.SDK.ParseJwtToken(tokenString)
		if err != nil {
			s.Logger.Error("Invalid token",
				zap.String("org", orgName),
				zap.Error(err))
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
			return
		}

		// Set context for downstream
		//c.Set("organization", orgName)
		// Set context for downstream
		ctx := utils.SetOrg(c.Request.Context(), orgName)
		c.Request = c.Request.WithContext(ctx)

		c.Set("user", verifiedClaims)
		c.Next()
	}
}

// GetOrg retrieves organization SDK by name
func (s *Security) GetOrg(name string) (*CasdoorOrg, error) {
	org, exists := s.Organizations[name]
	if !exists {
		return nil, fmt.Errorf("organization not found: %s", name)
	}
	return org, nil
}

func extractToken(c *gin.Context) (string, error) {
	h := c.GetHeader("Authorization")
	if h == "" {
		return "", fmt.Errorf("missing authorization header")
	}
	return strings.TrimPrefix(h, "Bearer "), nil
}

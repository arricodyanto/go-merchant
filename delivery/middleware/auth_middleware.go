package middleware

import (
	"go-merchant/shared/model"
	"go-merchant/shared/service"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	RequireToken(roles ...string) gin.HandlerFunc
}

type authMiddleware struct {
	jwtService service.JwtService
}

type AuthHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func (a *authMiddleware) RequireToken(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var autHeader AuthHeader
		if err := c.ShouldBindHeader(&autHeader); err != nil {
			log.Printf("RequireToken.autHeader: %v \n", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, &model.Status{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			})
			return
		}

		tokenHeader := strings.Replace(autHeader.AuthorizationHeader, "Bearer ", "", -1)
		if tokenHeader == "" {
			log.Printf("RequireToken.tokenHeader \n")
			c.AbortWithStatusJSON(http.StatusUnauthorized, &model.Status{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			})
			return
		}

		claims, err := a.jwtService.ParseToken(tokenHeader)
		if err != nil {
			log.Printf("RequireToken.ParseToken: %v \n", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, &model.Status{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			})
			return
		}
		c.Set("user", claims["username"])

		c.Next()
	}
}

func NewAuthMiddleware(jwtService service.JwtService) AuthMiddleware {
	return &authMiddleware{jwtService: jwtService}
}

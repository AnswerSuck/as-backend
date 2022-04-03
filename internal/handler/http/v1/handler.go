package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/answersuck/vault/internal/config"
	"github.com/answersuck/vault/internal/service"

	"github.com/answersuck/vault/pkg/auth"
	"github.com/answersuck/vault/pkg/logging"
	"github.com/answersuck/vault/pkg/validation"
)

const route = "/api/v1"

type Deps struct {
	Config          *config.Aggregate
	Logger          logging.Logger
	ErrorTranslator validation.ErrorTranslator
	TokenManager    auth.TokenManager
	AccountService  service.Account
	SessionService  service.Session
	AuthService     service.Auth
}

func SetupHandlers(e *gin.Engine, d *Deps) {
	// Options
	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	// Swagger UI
	e.Static(fmt.Sprintf("%s/swagger/", route), "third_party/swaggerui")

	e.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Resource handlers
	h := e.Group(route)
	{
		newSessionHandler(h, d)
		newAccountHandler(h, d)
		newAuthHandler(h, d)
	}
}

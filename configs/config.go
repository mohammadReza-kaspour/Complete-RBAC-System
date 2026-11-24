package configs

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/mohammadReza-kaspour/Complete-RBAC-System/pkg/ResultPattern"
)

const (
	//Db
	Host     = "localhost"
	User     = "postgres"
	Password = "admin"
	Dbname   = "projects"
	Port     = "5432"
	Schema   = "rbac"
	Sslmode  = "disable"
	TimeZone = "Asia/Tehran"

	//Auth
	JwtSecret = "mohammad"
)

type Config struct {
	FiberConfig *fiber.Config
}

func Load() *Config {
	return &Config{
		FiberConfig: fiberConfig,
	}
}

var fiberConfig = &fiber.Config{
	ErrorHandler: func(c fiber.Ctx, err error) error {
		// Check if it's a Api error
		if apiError, ok := err.(*ResultPattern.ErrorBase[map[string]any]); ok {
			return c.Status(apiError.StatusCode).JSON(apiError)
		}

		// Check if it's a Fiber error
		if fiberError, ok := err.(*fiber.Error); ok {
			apiError := ResultPattern.NewError(fiberError.Code, 9500, fiberError.Message, &map[string]any{})
			return c.Status(apiError.StatusCode).JSON(apiError)
		}

		apiError := ResultPattern.NewError(http.StatusInternalServerError, 9999, err.Error(), &map[string]any{})
		return c.Status(apiError.StatusCode).JSON(apiError)
	},
}
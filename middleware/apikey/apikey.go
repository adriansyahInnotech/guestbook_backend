package apikey

import (
	"fmt"
	"guestbook_backend/helper"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type ApiKey struct {
	helper *helper.Helper
}

func NewApiKey(helper *helper.Helper) *ApiKey {
	return &ApiKey{
		helper: helper,
	}

}

func (s *ApiKey) CheckAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing Authorization header",
			})
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid Authorization format",
			})
		}

		token := parts[1]
		prefix, key, ok := s.helper.Utils.ApiKey.SplitAPIKey(token)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid API key format",
			})
		}

		fmt.Printf("prefix : %s", prefix)
		fmt.Printf("key : %s", key)

		validate := s.helper.Utils.ApiKey.ValidateApiKey(prefix, key)

		if validate == false {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid API key",
			})
		}

		fmt.Println("\n\n\n validate : ", validate)

		// Lanjut ke handler berikutnya
		return c.Next()
	}
}

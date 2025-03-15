package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/somphonee/go-fiber-api/internal/models"
)


// Protected ตรวจสอบความถูกต้องของ JWT token
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
		"message": "Invalid or expired token",
	})
}

// GenerateToken สร้าง JWT token สำหรับผู้ใช้
func GenerateToken(user models.User) (string, error) {
	// สร้าง token ที่มีอายุ 24 ชั่วโมง
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// ลงนาม token ด้วย secret key
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// GetUserID ดึง user ID จาก token
func GetUserID(c *fiber.Ctx) uint {
	// ดึง authorization header
	authHeader := c.Get("Authorization")
	// ตัด Bearer prefix
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	// ตรวจสอบ token
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	// ดึง claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := uint(claims["user_id"].(float64))
		return userId
	}

	return 0
}
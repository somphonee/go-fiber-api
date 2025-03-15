package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// ErrorResponse แสดงข้อผิดพลาดของการ validate
type ErrorResponse struct {
	FailedField string `json:"field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
	Message     string `json:"message"`
}

// ValidateStruct ตรวจสอบโครงสร้าง struct ด้วย validator
func ValidateStruct(data interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()

	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = FormatErrorField(err.Field())
			element.Tag = err.Tag()
			element.Value = err.Param()

			switch err.Tag() {
			case "required":
				element.Message = element.FailedField + " is required"
			case "email":
				element.Message = element.FailedField + " must be a valid email"
			case "min":
				element.Message = element.FailedField + " must be at least " + err.Param() + " characters long"
			case "max":
				element.Message = element.FailedField + " must not be longer than " + err.Param() + " characters"
			default:
				element.Message = element.FailedField + " failed " + err.Tag() + " validation"
			}

			errors = append(errors, &element)
		}
	}

	return errors
}

// FormatErrorField แปลงชื่อฟิลด์ให้เป็นตัวพิมพ์เล็ก
func FormatErrorField(field string) string {
	// แปลงตัวแรกให้เป็นตัวพิมพ์เล็ก
	return strings.ToLower(field[:1]) + field[1:]
}

// HandleValidationErrors จัดการกับ error จากการ validate
func HandleValidationErrors(c *fiber.Ctx, errors []*ErrorResponse) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error":   "Validation failed",
		"details": errors,
	})
}
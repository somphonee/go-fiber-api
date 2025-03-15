package utils


import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Pagination ข้อมูลการแบ่งหน้า
type Pagination struct {
	Limit      int         `json:"limit,omitempty"`
	Page       int         `json:"page,omitempty"`
	TotalRows  int64       `json:"totalRows"`
	TotalPages int         `json:"totalPages"`
	Rows       interface{} `json:"rows"`
}

// GetPagination ดึงค่า pagination จาก request
func GetPagination(c *fiber.Ctx) (int, int) {
	// ค่าเริ่มต้น
	limit := 10
	page := 1

	// ดึงค่าจาก query parameters
	limitParam := c.Query("limit")
	pageParam := c.Query("page")

	// แปลงค่า limit
	if limitParam != "" {
		l, err := strconv.Atoi(limitParam)
		if err == nil && l > 0 {
			limit = l
		}
	}

	// แปลงค่า page
	if pageParam != "" {
		p, err := strconv.Atoi(pageParam)
		if err == nil && p > 0 {
			page = p
		}
	}

	return limit, page
}

// Paginate ทำ pagination กับ query
func Paginate(db *gorm.DB, result interface{}, pagination *Pagination) error {
	// คัดลอก DB เพื่อนำไปทำ count
	var totalRows int64
	db.Model(result).Count(&totalRows)

	// คำนวณ offset
	offset := (pagination.Page - 1) * pagination.Limit

	// ดึงข้อมูลตาม pagination
	err := db.Limit(pagination.Limit).Offset(offset).Find(result).Error

	// คำนวณจำนวนหน้าทั้งหมด
	pagination.TotalRows = totalRows
	pagination.Rows = result

	// คำนวณจำนวนหน้าทั้งหมด
	pagination.TotalPages = int((totalRows + int64(pagination.Limit) - 1) / int64(pagination.Limit))

	return err
}
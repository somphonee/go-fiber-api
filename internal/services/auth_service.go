package services

import (
	"errors"
	"github.com/somphonee/go-fiber-api/internal/repository"
	"github.com/somphonee/go-fiber-api/internal/models"
	"github.com/somphonee/go-fiber-api/internal/middleware"
	"golang.org/x/crypto/bcrypt"
	)

type AuthService struct {
userRepo *repository.UserRepository
}	

type RegisterInput struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}
func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}	


func (s *AuthService) Register(input RegisterInput) (*AuthResponse, error) {
	// ตรวจสอบว่ามีอีเมลนี้ในระบบหรือไม่
	_, err := s.userRepo.FindByEmail(input.Email)
	if err == nil {
		return nil, errors.New("email already exists")
	}

	// สร้างผู้ใช้ใหม่
	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	// บันทึกผู้ใช้ลงฐานข้อมูล
	if err := s.userRepo.Create(&user); err != nil {
		return nil, err
	}

	// สร้าง token
	token, err := middleware.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

func (s *AuthService) Login(input LoginInput) (*AuthResponse, error) {
	// ค้นหาผู้ใช้จากอีเมล
	user, err := s.userRepo.FindByEmail(input.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// ตรวจสอบรหัสผ่าน
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// สร้าง token
	token, err := middleware.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  user,
	}, nil
}
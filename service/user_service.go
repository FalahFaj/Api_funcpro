package service

import (
	"database/sql"
	"errors"
	"projek_funcpro_kel12/model"
	"projek_funcpro_kel12/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RegisterInput struct {
	Nama     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWT struct {
	UserId int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type UserService interface {
	Register(input RegisterInput) (*model.User, error)
	Login(input LoginInput) (string, error)
	GetUserById(id int64) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
	jwtToken string
}

func NewUserService(userRepo repository.UserRepository, jwtToken string) *userService {
	return &userService{userRepo, jwtToken}
}

func (s *userService) Register(input RegisterInput) (*model.User, error) {
	if input.Nama == "" || input.Email == "" || input.Password == "" || input.Role == "" {
		return nil, errors.New("tidak boleh ada yang kosong")
	}

	if input.Role != "pembeli" && input.Role != "petani" {
		return nil, errors.New("role harus pembeli atau petani")
	}

	sudahTerdaftar, err := s.userRepo.GetUserByEmail(input.Email)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	}

	if sudahTerdaftar != nil {
		return nil, errors.New("email sudah terdaftar")
	}

	user := &model.User{
		Nama:     input.Nama,
		Email:    input.Email,
		Password: input.Password,
		Role:     input.Role,
	}

	id, err := s.userRepo.Buat(user)
	if err != nil {
		return nil, err
	}

	user.Id = id

	return user, nil

}

func (s *userService) Login(input LoginInput) (string, error) {
	if input.Email == "" || input.Password == "" {
		return "", errors.New("tidak boleh ada yang kosong")
	}
	user, err := s.userRepo.GetUserByEmail(input.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("email tidak ditemukan")
		}
		return "", err
	}

	if input.Password != user.Password {
		return "", errors.New("password salah")
	}

	klaim := JWT{
		UserId: user.Id,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "funcpro_kel12",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, klaim)
	tokenString, err := token.SignedString([]byte(s.jwtToken))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *userService) GetUserById(id int64) (*model.User, error) {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user tidak ditemukan")
		}
		return nil, err
	}
	return user, nil
}

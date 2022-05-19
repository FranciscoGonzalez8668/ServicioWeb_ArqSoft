package services

import (
	userCliente "pan/clients/user"
	"pan/dto"
	"pan/model"
	e "pan/utils/errors"

	"github.com/golang-jwt/jwt"
)

type userService struct{}

type userServiceInterface interface {
	LoginUser(loginDto dto.LoginDto) (dto.Token, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

var jwtKey = []byte("secret_key")

func (s *userService) LoginUser(loginDto dto.LoginDto) (dto.Token, e.ApiError) {
	var user model.User = userCliente.GetUserByEmail(loginDto)

	var tokenDto dto.Token

	if user.Id_user == 0 {
		return tokenDto, e.NewBadRequestApiError("User not found")
	}

	if user.Password == loginDto.Password {
		token := jwt.New(jwt.SigningMethodHS256)
		tokenString, _ := token.SignedString(jwtKey)
		tokenDto.Token = tokenString
	}
	return tokenDto, nil
}

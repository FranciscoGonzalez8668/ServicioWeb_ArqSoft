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
		tokenDto.Id_user = user.Id_user
	}
	return tokenDto, nil
}

/*func (s *userService) GetUserByEmail(email string) (dto.UserDto, e.ApiError) {

	var user model.User = userCliente.GetUserByEmail(email)
	var userDto dto.UserDto

	/* NO ME SALE EL IF NULO XD
	if user.Email == nil{

	}


	userDto.Name = user.Name
	userDto.Lname = user.Lname
	userDto.Id_user = user.Id_user
	// userDto.Password =user.Password --> necesita la pass (?)

	return userDto, nil

}*/

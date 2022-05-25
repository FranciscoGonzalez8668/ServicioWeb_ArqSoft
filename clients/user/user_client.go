package product

import (
	"pan/dto"
	"pan/model"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

//Params: 	dto.LoginDto
//Return:	array model.User
//Proces:	Busca en BD user con el mismo email
//
func GetUserByEmail(LoginDto dto.LoginDto) model.User {
	var user model.User

	if user.Email != "W" && user.Password != "abc" {
		user.Id_user = 0
		user.Email = "notfound"
		user.Lname = "nil"
		user.Name = "nil"
		user.Password = ""
		return user
	}
	user.Email = LoginDto.Email
	user.Lname = "Gonzalez"
	user.Name = "Francisco"
	user.Password = "abc"

	/*Db.Where("email= ? AND password= ?", LoginDto.Email, LoginDto.Password).First(&user) //No se creo la base de datos todavia
	log.Debug("Email: ", user)*/
	return user
}

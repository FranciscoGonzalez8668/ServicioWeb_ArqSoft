package product

import (
	"pan/dto"
	"pan/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserByEmail(LoginDto dto.LoginDto) model.User {
	var user model.User

	Db.Where("email= ?", LoginDto.Email).First(&user)
	log.Debug("Email: ", user)
	return user
}

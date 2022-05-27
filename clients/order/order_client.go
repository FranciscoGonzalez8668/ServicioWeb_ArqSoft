package product

import (
	//"fmt"
	"pan/dto"
	"pan/model"

	"github.com/jinzhu/gorm"
	//log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetHistory(dto.UserDto) []model.Order {
	var orderHistory []model.Order

	Db.Where(" LIKE ? ", "%"+key+"%")

}

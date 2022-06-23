package sql

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"muback/config"
)

type Admin  struct {
	Username string `db:"username"`
	Password string `db:"password"`
	Bind string `db:"bind"`
	Role string `db:"role"`
}

type User struct {
	Uid int `db:"uid"`
	Pid int `db:"pid"`
	Address string `db:"address"`
	OutedU float32 `db:"outed_u"`
	OutableU float32 `db:"outable_u"`
	TotalU float32 `db:"total_u"`
	OutsurplusU float32 `db:"outsurplus_u"`
	TodayU float32 `db:"today_u"`
	Fund int `db:"fund"`
}

var DB *gorm.DB
func InitDatabase()  {
	sqlSource := config.Cfg.Database.DBUserName + ":" +config.Cfg.Database.DBPassword + "@tcp(" + config.Cfg.Database.DBHost + ":" + config.Cfg.Database.DBPort + ")/"+config.Cfg.Database.DBSchema+"?"+config.Cfg.Database.DBArgs
	db, err := gorm.Open("mysql", sqlSource)
	if err !=nil {
		log.Println(err)
	}
	db.DB().SetMaxIdleConns(10)
	DB =db
}





package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"mubackend/config"
	"time"
)

type User struct {
	Uid int `db:"uid"`
	Address string `db:"address"`
	Pid int `db:"pid"`
	Outed int64 `db:"outed"`
	Outable int64 `db:"outable"`
	Outsurplus int64 `db:"outsurplus"`
	Total int64 `db:"total"`
	Treasure int `db:"treasure"`
	Fund int `db:"fund"`
	TodayAmount int64 `db:"today_amount"`
	OutedU	float32 `db:"outed_u"`
	OutableU float32 `db:"outable_u"`
	TotalU float32 `db:"total_u"`
	OutsurplusU float32 `db:"outsurplus_u"`
	TodayU	float32 `db:"today_u"`
	TodayFund int64 `db:"today_fund""`
	OutableFund int64 `db:"outable_fund"`
	Authority int `db:"authority"`
}
type Order struct {
	Oid int `db:"oid"`
	Uid int `db:"uid"`
	Inub int `db:"inub"` //1是USDT，2是BNB
	Tradehash string `db:"tradehash"`
	CreateTime time.Time `db:"create_time"`
	Type int `db:"type"` //1为treasure，2为理财，3为基金
}

type Article struct {
	Title string `db:"tile"`
	Content string `db:"content"`
	Value string `db:"value"`
}

var DB *gorm.DB

func InitDatabase()  {
	sqlSource := config.Cfg.Database.DBUserName + ":" +config.Cfg.Database.DBPassword + "@tcp(" + config.Cfg.Database.DBHost + ":" + config.Cfg.Database.DBPort + ")/"+config.Cfg.Database.DBSchema+"?"+config.Cfg.Database.DBArgs
	db, err := gorm.Open("mysql", sqlSource)
	if err !=nil {
		log.Fatal(err)
	}
	db.DB().SetMaxIdleConns(10)
	DB =db
}

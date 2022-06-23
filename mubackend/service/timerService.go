package service

import (
	"github.com/robfig/cron/v3"
	"log"
	"mubackend/sql"
)

func InitCron()  {
	crontab := cron.New(cron.WithSeconds())
	log.Println("timer cron start ...")
	spec :="0 0 1 * * ?"
	crontab.AddFunc(spec,Today)
	crontab.Start()
}
//把，today_amount置为0，分配todayamount和fund
func Today()  {
	var rate float32
	result, err := GetRate()

	if err!=nil {
		result = float32(0.51)
	}
	switch int(result/0.1){
	case 5:
		rate=float32(0.005)
	case 4:
		rate=float32(0.004)
	case 3:
		rate=float32(0.003)
	case 2:
		rate=float32(0.002)
	case 1,0:
		rate=float32(0.001)
	default:
		rate=float32(0.006)
	}

	//sql.UpdateTodayAmountZero()

	yestodayAmount := sql.GetFundAmount()
	amount, err := GetFundAmount()
	if err!=nil {
		amount = yestodayAmount.Fundamount
	}
	UpdateTodayAmountEveryDay(rate,amount-yestodayAmount.Fundamount)
	sql.GetUserInfoAll()
	UpdateFinalSon()
	//让因为团队奖励分配完成，让算力降为0
	sql.UpdateCalZero()
	sql.UpdateFundAmount(amount)
}

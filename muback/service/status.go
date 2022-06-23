package service

import "muback/sql"

func GetStatus() (bool,bool) {
	noticeStatus := sql.NoticeStatus()
	trade, _ := GetStatusAutoTrade()

	return noticeStatus,trade
}

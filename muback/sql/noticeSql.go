package sql

func CloseNotice()  {
	DB.Exec("update content set value =0 where tid in (7,8)")
}

func ChangePrice(price int)  {
	DB.Exec("update content set value=? where tid=9", price)

}


func OpenNotice(openNotice bool)  {
	if openNotice {
		DB.Exec("update content set value =1 where tid in (7,8)")
	}else {
		DB.Exec("update content set value =0 where tid in (7,8)")
	}
}

func ReleaseNotice(chineseText string,englishText string)  {
	DB.Exec("update content set content =? where tid=7",chineseText)
	DB.Exec("update content set content =? where tid=8",englishText)
	OpenNotice(true)
}

type value struct {
	Value int `db:"value"`
}
func NoticeStatus()bool  {
	var va value
	DB.Raw("select value from content where tid =7").Scan(&va)
	if va.Value>0 {
return true
	}
	return false
}

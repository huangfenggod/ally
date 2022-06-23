package sql

import (
	"log"
	"time"
)

type FundAmount struct {
	Fundamount int64 `db:"fundamount"`
}
//<<<<<<<<<<<<<<<<查询操作>>>>>>>>>>>>>>>>
//通过address查询用户信息
func GetInfoByAddress(address string) User{
	var user User
	DB.Raw("select uid,address ,treasure ,fund ,outed_u ,outable_u ,today_u,total_u ,outsurplus_u ,today_fund,outable_fund ,authority from user where address =? and exists(select 1 from user where address =? limit 1) ", address,address).Scan(&user)
	return user
}
type TeamTotal struct {
	Team float32 `db:"team"`
}

//查询某个id团队下全部业绩
func GetTeamTotalU(uid int)float32{
	var tt TeamTotal
	DB.Raw("with recursive cte  as(select uid ,pid ,total_u from user where pid =? union all select u.uid ,u.pid ,u.total_u from cte c inner join user u on c.uid =u.pid) select sum(total_u) team from cte ",uid).Scan(&tt)
	return tt.Team
}



//查取昨日fundamount
func GetFundAmount()FundAmount  {
	var fundamount FundAmount
	DB.Debug().Raw("select * from fund where id =1").Scan(&fundamount)
	return fundamount
}
//修改今日fundamount
func UpdateFundAmount(amount int64)  {
	DB.Exec("update fund set fundamount = ?",amount)
	log.Printf("change fundamount :%d",amount)
}

func GetOutAble(address string) User {
	var user  User
	DB.Debug().Raw("select uid ,outable_u, outable_fund from user where address =?",address).Scan(&user)
	return user
}
func ExistsIdByAddress(address string) int {
	var id []User
	DB.Raw("select uid from user where address = ? and exists (select 1 from user where address = ? limit 1)" ,address,address).Scan(&id)
	if  len(id)!=0{
		return id[0].Uid
	}else {
		return 0
	}
}
//查询目前已经领取多少treasure
func GetNumberOfTreasure() int {
	var count int
	DB.Table("user").Where("treasure =1 or treasure =2").Count(&count)
	return count
}

//查询出整个outsurplus大于0和fund>0的user表
func GetUserOutsurplus() []User  {
	var users []User
	DB.Raw("select uid ,pid ,outsurplus ,total ,outsurplus_u ,total_u from user where outsurplus >0 or fund >0").Scan(&users)
	return users
}

//查取文章
func GetArticle(id int)Article  {
	var article Article
	DB.Raw("select * from content where tid =?",id).Scan(&article)
	return article
}

//根据address查询订单情况
func GetOrderInfo(address string)[]Order{
	var orders []Order
	DB.Raw("select B.* from user A, oder B where A.address =? ",address).Scan(&orders)
	return orders
}
//查询出所有信息
func GetUserInfoAll() ([]User ,error) {
	var  users []User
	sel := DB.Raw("select uid ,pid ,total,total_u,outable ,outsurplus,today_amount,outable_u,outsurplus_u,today_u from user ").Scan(&users)

	if sel.Error !=nil {
		log.Println(sel.Error)
		return users,sel.Error
	}
	return users,nil
}

type cou struct {
	Uid int `db:"uid"`
}
func GetTeamNumber(address string)int {
	//var allCount int
	//DB.Table("user").Count(&allCount)
	id := ExistsIdByAddress(address)
	if id ==0 {
		return 0
	}
	var co []cou
	DB.Raw("select uid from user where pid =?",id).Scan(&co)
	return len(co)

}

type countNumber struct {
	Uid int `db:"uid"`
}

func GetSixTeamNumber(address string)int {
	uid := ExistsIdByAddress(address)
	if uid ==0 {
		return 0
	}
	sum := 1
	number := GetSonTeamNumber(uid)
	sum += len(number)
	for a := 0; a < len(number); a++ {
		number1 := GetSonTeamNumber(number[a].Uid)
		sum += len(number1)
		for b := 0; b < len(number1); b++ {
			number2 := GetSonTeamNumber(number1[b].Uid)
			sum += len(number2)
			for c := 0; c < len(number2); c++ {
				number3 := GetSonTeamNumber(number2[c].Uid)
				sum += len(number3)
				for d := 0; d < len(number3); d++ {
					number4 := GetSonTeamNumber(number3[d].Uid)
					sum += len(number4)
					for e := 0; e < len(number4); d++ {
						number5 := GetSonTeamNumber(number4[e].Uid)
						sum += len(number5)
						for f := 0; f < len(number5); f++ {
							sum += len(number5)
						}
					}
				}
			}
		}


	}
	return sum
}

type Num struct {
	Number int `db:"number"`
}
func GetAllTeamNumber(address string) int {
	id := ExistsIdByAddress(address)
	if id ==0 {
		return 0
	}
	var num Num
	DB.Raw("with recursive cte  as(select uid ,pid from user where pid =? union all select u.uid ,u.pid  from cte c inner join user u on c.uid =u.pid) select count(uid) number from cte ",id).Scan(&num)
	return num.Number
}



func GetPriceOfContent()int {
	var num Num
	DB.Raw("select value number from content where tid=9").Scan(&num)
	return num.Number
}
func ChangePriceContent(price int)  {
	DB.Exec("update content set value =? where tid=9",price)
}





func GetSonTeamNumber(uid int)[]countNumber{
	var countN []countNumber
	if uid ==0{
		return countN
	}
	DB.Raw("select * from user where pid=? ",uid).Scan(&countN)
	return countN
}





//<<<<<<<<<<<<<<<<插入操作>>>>>>>>>>>>>>>>

func InsertUserByTreasureAndPid(address string ,pid int) {
	exec := DB.Debug().Exec("insert into  user(address,pid,treasure) values(?,?,1)",address,pid)
	if exec.Error !=nil {
		log.Println(exec.Error)
	}

	log.Printf("crate user address:%s pid:%d treasure\n", address, pid)
}

func InsertUserByTreasure(address string)  {
	exec := DB.Exec("insert into  user(address ,treasure) values(?,1)",address)
	if exec.Error !=nil {
		log.Println(exec.Error)
	}
	log.Printf("crate user address:%s by treasure\n",address)
}

func InsertUserByFundAndPid(address string ,pid int,fund int)  {
	exec := DB.Exec("insert into  user(address ,pid,fund) values(?,?,?)",address,pid,fund)
	if  exec.Error !=nil{
		log.Println(exec.Error)
	}
	log.Printf("crate user address:%s pid:%d by fund\n", address, pid)
}

func InsertUserByFund(address string ,fund int)  {
	exec := DB.Exec("insert into  user(address ,fund) values(?,?)",address,fund)
	if exec.Error !=nil {
		log.Println(exec.Error)
	}
	log.Printf("crate user address:%s by fund\n", address )
}

func InsertUserByInvestAndPid(address string,pid int )  {
	exec := DB.Exec("insert into  user(address ,pid) values(?,?)",address,pid)
	if exec.Error !=nil {
		log.Println(exec.Error)
	}
	log.Printf("crate user address:%s pid:%d by invest\n", address,pid )
}

func InsertUserByInvest(address string)  {
	exec := DB.Exec("insert into user(address ) values(?)",address)
	if exec.Error !=nil {
		log.Println(exec.Error)
	}
	log.Printf("crate user address:%s by invest\n", address )
}
func InsertOrder(uid int , inAmount float32, tradeHash string ,types int)  {
	exec := DB.Exec("insert into oder (uid, inub ,tradehash ,create_time, type) values(?,?,?,?,?) ",uid,inAmount,tradeHash,time.Now(),types)
	if exec.Error !=nil{
		log.Println(exec.Error)
	}
	log.Printf("crate order uid:%d amount:%f trade:%s type:%d by invest\n", uid,inAmount,tradeHash,types )
}

//<<<<<<<<<<<<<<<<更新user>>>>>>>>>>>>>>>>

func UpdateUserFund(address string, fund int )  {
	exec := DB.Exec("update user set treasure =2, fund = fund +? where address =?",fund,address)
	if exec.Error !=nil {
		log.Println(exec.Error)
	}
	log.Printf("update address:%s fund increase:%d \n",address,fund)
}

func GetTreasureByAddress(address string)  {
	DB.Exec("update user set treasure  =1 where treasure = 0 and address =?",address)
	log.Printf("update user address:%s acquire treasure",address)
}

func UpdateUserTreasure(address string) (bool,error)  {
	exec := DB.Exec("update user set treasure = 2 where address =?",address)
	if exec.Error !=nil {
		log.Println(exec.Error)
		return false,exec.Error
	}
	log.Printf("update address:%s use treasure\n", address)
	return true,nil
}

func UpdateUserByInvest(address string,amountU float32,amountGetU float32,amount int64,amountGet int64) (bool,error) {
	exec := DB.Exec("update user set total_u = total_u + ?, outsurplus_u = outsurplus_u + ?, total = total + ?,  outsurplus = outsurplus + ? where address =?",amountGetU,amountGetU,amountGet,amountGet,address )
	if exec.Error !=nil {
		log.Println(exec.Error)
		return false,exec.Error
	}
	log.Printf("update user address:%s payu:%f getu:%f token:%d gettoken:%d ",address,amountU,amountGetU,amount,amountGet)
	return true,nil
}

func UpdateTodayAmountByAddress(address string,amountU float32 ,amount int64)  {
	exec := DB.Exec("update user set today_u = if(? > outsurplus_u ,outsurplus ,?),outable_u = outable_u + today_u, outsurplus_u = outsurplus - today_u,today_amount = if(? > outsurplus ,outsurplus ,? ),outable = outable - today_amount, outsurplus = outsurplus - today_amount where address =?",amountU,amountU,amount,amount,address)
	if exec.Error !=nil {
		log.Println(exec.Error)
	}
	log.Printf("update today_amount address:%s today_u:%f today_amount:%d",address,amountU,amount)
}

func UpdateTodayAmountByUid(uid int,amountU float32 ,amount int64)  {
	exec := DB.Exec("update user set today_u = if(? > outsurplus_u ,outsurplus ,?),outable_u = outable_u + today_u, outsurplus_u = outsurplus - today_u,today_amount = if(? > outsurplus ,outsurplus ,? ),outable = outable - today_amount, outsurplus = outsurplus - today_amount where uid =?",amountU,amountU,amount,amount,uid)
	if exec.Error !=nil {
		log.Println(exec.Error)
	}
	log.Printf("update today_amount uid:%d today_u:%f today_amount:%d",uid,amountU,amount)
}

func UpdateTodayAmounAndFundtByUid(uid int,amountU float32 ,amount int64,todayFund int64)  {
	exec := DB.Exec("update user set today_u = if(? > outsurplus_u ,outsurplus ,?),outable_u = outable_u + today_u, outsurplus_u = outsurplus - today_u,today_amount = if(? > outsurplus ,outsurplus ,? ),outable = outable - today_amount, outsurplus = outsurplus - today_amount ,today_fund =? ,outable_fund = outable_fund + ?where uid =?",amountU,amountU,amount,amount,uid,todayFund,todayFund)
	if exec.Error !=nil {
		log.Println(exec.Error)
	}
	log.Printf("update today_amount uid:%d today_u:%f today_amount:%d todayfund:%d",uid,amountU,amount,todayFund)
}

func UpdateCash(address string)  {
	exec:= DB.Exec("update user set outed = outed + outable, outable = 0, outed_u = outed_u + outable_u , outable_u = 0,outable_fund =0 where address =?",address)
	if exec.Error !=nil {
		log.Println(exec.Error)
	}
	log.Printf("update address: %s cash",address)
}
//每日更新所有
func UpdateAllEveryDay(rate float32 ,fundGet int64 )  {
	exec := DB.Debug().Exec("update user u  inner join (select sum(fund) sum from user) temp set u.today_amount = if(u.total * ? > u.outsurplus,u.outsurplus,u.total * ?) ,u.today_u = if(u.total_u * ? > u.outsurplus_u,u.outsurplus_u,u.total_u * ?), u.today_fund = if(u.fund>0,(u.fund / temp.sum)*?,0),u.outable_fund = u.outable_fund + if(u.fund>0,?*u.fund / temp.sum,0) ,u.outable = u.outable + if(u.total *?  > u.outsurplus,u.outsurplus,u.total *? ),u.outsurplus = u.outsurplus - if(u.total *?  > u.outsurplus,u.outsurplus,u.total *? ),u.outable_u = u.outable_u + if(u.total_u *? > u.outsurplus_u,u.outsurplus_u,u.total_u *? ),u.outsurplus_u=u.outsurplus_u - if(u.total_u *? > u.outsurplus_u,u.outsurplus_u,u.total_u *? )", rate, rate, rate, rate, fundGet, fundGet, rate, rate, rate, rate, rate, rate, rate, rate)
	if exec.Error !=nil {
		log.Println(exec.Error)
	}
	DB.Exec("update  user set outable = outsurplus + outable,outsurplus=0,outable_u  = outable_u + outsurplus_u,outsurplus_u=0 where outsurplus_u <today_u")
	DB.Debug().Exec("update user set total_u =0 ,total=0 where outsurplus_u = 0")

	log.Println("update everday todayamount...")
}

func UpdateCalZero()  {
	DB.Exec("update user set total_u=0,total=0 where outsurplus_u=0")
}

func UpdateTodayAmountZero()  {
	DB.Debug().Exec("update user set today_amount = 0,today_u =0 ,today_fund =0")
}
//使用treasure增加收益
func UpdateTreasureAndIncome(addres string)(bool,error){
	exec := DB.Exec("update user set total=outsurplus*0.5 + total , outsurplus= outsurplus * 1.5,total_u = outsurplus_u *0.5 + total_u,outsurplus_u = outsurplus_u *1.5 ,treasure =2 where address =?", addres)
	if exec.Error !=nil {
		log.Println(exec.Error)
		return false,exec.Error
	}
	return true,nil
}













package service

import (
	"errors"
	"log"
	"mubackend/sql"
)

type ResponseInfo struct {
	Uid int	`json:"uid"`
	Address string `json:"address"`
	TodayU float32	`json:"todayU"`
	TotalU float32	`json:"totalU"`
	Treasure int	`json:"treasure"`
	Fund int	`json:"fund"`
	OutedU float32	`json:"outedU"`
	OutAbleU float32	`json:"OutAbleU"`
}

func GetContent(tid int) sql.Article {
	return sql.GetArticle(tid)
}
//获取用户信息
func GetUserInfo(address string ,price float32) ResponseInfo{
	info := sql.GetInfoByAddress(address)
	fund := float32(info.TodayFund)/(price*100000000)
	outAbleU := float32(info.OutableFund)/(price*100000000)
	responseInfo := ResponseInfo{
		Uid: info.Uid,
		Address: info.Address,
		TodayU: info.TodayU + fund,
		Treasure: info.Treasure,
		TotalU: info.TotalU,
		Fund: info.Fund,
		OutAbleU: info.OutableU + outAbleU,
		OutedU: info.OutedU,
	}
	return responseInfo
}
//
func GetTreasureNumber()int {
	treasure := sql.GetNumberOfTreasure()
	return treasure
}

func GetTeamTotalU(address string)float32  {
	id := sql.ExistsIdByAddress(address)
	if id==0{
		return 0
	}
	totalU := sql.GetTeamTotalU(id)
	return totalU
}



//获取treasure
func AcquireTreasure(address string,tradeHash string ,pid int)(bool,error) {
	//这里验证BNB是否付款成功
	identify := IdentifyTransactionHash2(tradeHash)
	if !identify {
		return false ,errors.New("check pay invalid")
	}
	//验证该地址是否获取过treasure
	//获取获取宝物
	treasure := sql.GetNumberOfTreasure()
	if treasure>=10000 {
		log.Fatalf("acquire address:%s tresure more than 10000,false",address)
		return false ,errors.New("token more than 10000")
	}
	if sql.ExistsIdByAddress(address)>0{
		var info sql.User
		sql.DB.Raw("select treasure from user where address =?",address).Scan(&info)
		if info.Treasure>0 {
			log.Println("get token fail ,because already got a token")
			return false ,errors.New("get token fail ,because already got a token")
		}
		sql.GetTreasureByAddress(address)
		InevstTen(address)
		return true, nil
	}else {
		sql.InsertUserByTreasureAndPid(address,pid)
		InevstTen(address)
		return true,nil
	}
}
//使用宝物是后面收益增加50%

func UseTreasure(address string) (bool,error) {
	id := sql.ExistsIdByAddress(address)
	if id ==0{
		return false,errors.New("address not exist")
	}
	info := sql.GetInfoByAddress(address)
	if info.Treasure !=1 || info.TotalU==0 {
		return false ,errors.New("have no calculation power")
	}
	income, err := sql.UpdateTreasureAndIncome(address)
	if !income{
		return false ,err
	}
	return true,nil
}

func GetTeamNumber(address string) (int,int) {
	teamNumber :=  sql.GetTeamNumber(address)
	number := sql.GetAllTeamNumber(address)
	return number, teamNumber
}

//参与基金，必须要有treasure才行
func JoinFundWithTreasure(address string ,tradeHash string )(bool,error)  {
	//验证U tradehash的情况
	identify := IdentifyTransactionHash2(tradeHash)
	if !identify {
		return false ,errors.New("tradeHash wrong join fund fail")
	}
	id := sql.ExistsIdByAddress(address)
	if id ==0 {
		return false ,errors.New("this address dont exists")
	}
	user := sql.GetInfoByAddress(address)
	if user.Treasure !=1 || user.Fund !=0{
		return false ,errors.New("no treasure can use or has fund already")
	}
	log.Printf("update address:%s fund 1",address)
	sql.UpdateUserFund(address,1)
	sql.InsertOrder(user.Uid,100,tradeHash,2)
	return true,nil
}
//参与基金多份
func JoinFund(address string ,number int, tradeHash string ,amount float32 )(bool,error)  {
	//验证U tradehash的情况
	exists := sql.ExistsIdByAddress(address)
	if exists==0{
		log.Printf("buy fund fail because have no this address:%s",address)
		return false,errors.New("no user")
	}
	log.Fatalf("update address:%s fund :%d",address,number)
	sql.UpdateUserFund(address,number)
	info := sql.GetInfoByAddress(address)
	sql.InsertOrder(info.Uid,amount,tradeHash,number)
	return true,nil
}

//参与invest
func Invest(address string,tradeHash string , amountU float32,price float32,pid int)(bool,error) {
	//验证U tradehash的情况
	identify := IdentifyTransactionHash2(tradeHash)
	if !identify {
		return false ,errors.New("tradeHash wrong")
	}
	if amountU <50 {
		return false,errors.New("less than 50U")
	}
	if sql.ExistsIdByAddress(address) ==0{
		//创建用户
		sql.InsertUserByInvestAndPid(address,pid)
	}
	user := sql.GetInfoByAddress(address)
	sql.InsertOrder(user.Uid,amountU,tradeHash,1)
	amount := amountU  *100000000
	i := int64(amount)
	sql.UpdateUserByInvest(address,amountU,amountU*2,i,i*2)
	return true,nil
}

func InevstTen(address string) (bool,error) {

	user := sql.GetInfoByAddress(address)
	sql.InsertOrder(user.Uid,10,"0x111111111111111111111",1)
	amount := 10  *100000000
	i := int64(amount)
	sql.UpdateUserByInvest(address,0,10,0,i)
	return true,nil
}


//宝物体验invest,amountU =10
func InvestByTreasure(address string,tradeHash string , amountU float32,price float32)(bool,error)  {
	identify := IdentifyTransactionHash2(tradeHash)
	if !identify {
		return false ,errors.New("tradeHash wrong")
	}
	info := sql.GetInfoByAddress(address)
	if info.Treasure !=1{
	return false,errors.New("have no treasure can not take part in invest")
	}
	sql.InsertOrder(info.Uid,amountU,tradeHash,1)
	amount := amountU *100000000
	i := int64(amount)
	invest, err := sql.UpdateUserByInvest(address, amountU, amountU*5, i, i*5)
	if !invest {
		return false,err
	}
	treasure, err := sql.UpdateUserTreasure(address)
	if !treasure {
		return false,err
	}
	return true,nil
}

//发起转账交易
func CashOut(address string ,tradeHash string) (bool,error){
	identify := IdentifyTransactionHash2(tradeHash)
	if !identify {
		return false ,errors.New("tradeHash wrong ")
	}
	able := sql.GetOutAble(address)



	//fmt.Println(able.Outable)
	log.Printf("verify pass ,begin transfer to %s , %d",address,able.Outable)
	if able.OutableU+float32(able.OutableFund) <=0{
	return false ,errors.New("have not any amount to cash ")
	}
	//调取提取函数
	price := GetPrice()
	outable :=int64(price * able.OutableU * 100000000)
	txhash, err := Transfer(address, outable, able.OutableFund)
	if err!=nil {
		return false,err
	}
	sql.CashRecorder(txhash,address,outable+able.OutableFund,price)
	sql.UpdateCash(address)

	return true,nil
}


func UpdateTodayAmountEveryDay(rate float32,fund int64)  {
	sql.UpdateAllEveryDay(rate,fund)

}


//处理子代的情况
func UpdateFinalSon()  {
	users,err := sql.GetUserInfoAll()
	if err !=nil{
		log.Printf("exec son todayamount fail")
		return
	}
	for i :=0;i<len(users);i++ {
		//计算分代情况
		amount := getSixAmount(users[i].Uid, users)
		if users[i].Outsurplus >amount {
			users[i].Outsurplus -=amount
			users[i].Outable +=amount
			users[i].TodayAmount +=amount
			users[i].OutsurplusU = float32(users[i].Outsurplus) / float32(users[i].Total) *users[i].TotalU
			users[i].OutableU = float32(users[i].Outable) / float32(users[i].Total) *users[i].TotalU
			users[i].TodayU = float32(users[i].TodayAmount) / float32(users[i].Total) *users[i].TotalU

		}else {
			users[i].Outable += users[i].Outsurplus
			users[i].TodayAmount += users[i].Outsurplus
			users[i].Outsurplus =0
			users[i].OutableU +=users[i].OutsurplusU
			users[i].TodayU += users[i].OutsurplusU
			users[i].OutsurplusU =0
		}
		sql := sql.DB.Exec("update user set today_amount =?,outsurplus =?, outable=?,today_u=?,outsurplus_u=?, outable_u =? where uid =?",users[i].TodayAmount,users[i].Outsurplus,users[i].Outable,users[i].TodayU,users[i].OutsurplusU,users[i].OutableU,users[i].Uid)
		if sql.Error !=nil {
			log.Println(sql.Error)
			log.Fatalf("son update uid :%d fail",users[i].Uid)
		}
	}
}

func sonAmount(uid int,users []sql.User) (int64 ,[]int) {
	amount := int64(0)
	var ints []int
	for i :=0;i<len(users);i++ {
		if users[i].Pid == uid {
			ints = append(ints, users[i].Uid)
			amount += users[i].TodayAmount
		}
	}
	return amount ,ints
}

func getSixAmount(uid int,users []sql.User) int64{
	amount := int64(0)
	amount1, int1 := sonAmount(uid, users)
	amount += amount1 /2
	for i :=0;i<len(int1);i++ {
		amount2, int2 := sonAmount(int1[i], users)
		amount += amount2 *2/5
		for j:=0;j<len(int2);j++  {
			amount3, int3 := sonAmount(int2[j], users)
			amount += amount3 *3/10
			for m :=0;m<len(int3);m++ {
				amount4, int4 := sonAmount(int3[m], users)
				amount += amount4 /5
				for n :=0;n<len(int4);n++ {
					amount5, int5 := sonAmount(int4[n], users)
					amount += amount5 /10
					for k:=0;k<len(int5);k++{
						amount6, _ := sonAmount(int5[k], users)
						amount += amount6 /20
					}
				}
			}
		}
	}
	return amount
}






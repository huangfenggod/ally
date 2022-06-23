package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"mubackend/crypto"
	"mubackend/service"
	"mubackend/sql"
	"net/http"
	"strconv"
)

type Param struct {
	Address string `json:"address"`
	Password string `json:"password"`
	Value string `json:"value"`
	Paddress string `json:"paddress"`
	Amount float32 `json:"amount"`
}



func ApiGroup() {
	group := GIN.Group("/v1/api")
	{
		//调整价格
		group.GET("/cprice", func(context *gin.Context) {
			price := context.Query("price")
			if len(price)==0 {
				context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "no price allow"})
				return
			}
			f, err := strconv.ParseFloat(price, 32)
			if err!=nil {
				context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "no price allow"})
				return
			}
			//service.PriceControll= float32(float)
			sql.ChangePriceContent(int(f))
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "ok"})
		})

		//获取文案，标题广告语
		group.GET("/text", func(context *gin.Context) {

			language := context.GetHeader("accept-language")

			if language == "zh-CN" {
				content1 := service.GetContent(1)
				content3 := service.GetContent(3)
				content5 := service.GetContent(5)
				content7 := service.GetContent(7)
				context.JSON(http.StatusOK, ResponseUtil{Status: true, Data: gin.H{"first": content1, "second": content3, "third": content5,"notice":gin.H{"status":content7.Value,"content":content7.Content}}, Msg: "success"})
			} else {
				content2 := service.GetContent(2)
				content4 := service.GetContent(4)
				content6 := service.GetContent(6)
				content8 := service.GetContent(8)
				context.JSON(http.StatusOK, ResponseUtil{Status: true, Data: gin.H{"first": content2, "second": content4, "third": content6,"notice":gin.H{"status":content8.Value,"content":content8.Content}}, Msg: "success"})
			}

		})
		//获取用户信息
		group.GET("/user", func(context *gin.Context) {
			language := context.GetHeader("accept-language")
			address := context.Query("address")
			if language =="zh-CN" {
				if len(address) == 0 {
					context.JSON(http.StatusNotFound, ResponseUtil{Msg: "钱包地址无效", Status: false})
					return
				}
				price := service.GetPrice()
				if price == 0 {
					price = 1
				}
				info := service.GetUserInfo(address, price)
				context.JSON(http.StatusOK, ResponseUtil{Msg: "成功", Status: true, Data: gin.H{"Id": info.Uid, "IsHasTreasure": info.Treasure, "CashAble": info.OutAbleU, "TodayGet": info.TodayU, "IsHasFund": info.Fund, "Address": info.Address, "Total": info.TotalU, "Already": info.OutAbleU + info.OutedU}})
			}else {
				if len(address) == 0 {
					context.JSON(http.StatusNotFound, ResponseUtil{Msg: "address invalid", Status: false})
					return
				}
				price := service.GetPrice()
				if price == 0 {
					price = 1
				}
				info := service.GetUserInfo(address, price)
				context.JSON(http.StatusOK, ResponseUtil{Msg: "success", Status: true, Data: gin.H{"Id": info.Uid, "IsHasTreasure": info.Treasure, "CashAble": info.OutAbleU, "TodayGet": info.TodayU, "IsHasFund": info.Fund, "Address": info.Address, "Total": info.TotalU, "Already": info.OutAbleU + info.OutedU}})
			}
		})
		//获取令牌数量

		group.GET("/getTreasure", func(context *gin.Context) {
			language := context.GetHeader("accept-language")
			number := service.GetTreasureNumber()
			if language == "zh-CN"{
				context.JSON(http.StatusOK, ResponseUtil{Msg: "成功", Data: gin.H{"number": number}, Status: true})
				return
			}else {
				context.JSON(http.StatusOK, ResponseUtil{Msg: "success", Data: gin.H{"number": number}, Status: true})
				return
			}
		})
		//获取令牌
		group.POST("/treasure", func(context *gin.Context) {
			language := context.GetHeader("accept-language")
			address := context.Query("address")
			tradeHashEncode := context.Query("password")
			paddress := context.Query("paddress")
			if len(address) == 0 && len(tradeHashEncode) == 0{
				var param Param
				context.BindJSON(&param)
				address = param.Address
				tradeHashEncode = param.Password
				paddress = param.Paddress
			}
				tradeHash, err2 := crypto.BaseDecode(tradeHashEncode)
			if language =="zh-CN" {
				if err2 != nil {
					context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "hash值错误"})
					return
				}

				if len(address) == 0 || len(tradeHash) == 0 {
					context.JSON(http.StatusOK, ResponseUtil{Msg: "没有这个地址或hash", Status: false})
					return
				}
				pid := 0
				if len(paddress) > 0 {
					info := sql.ExistsIdByAddress(paddress)
					if info == 0 {
						context.JSON(http.StatusOK, ResponseUtil{Msg: "推荐人地址无效", Status: false})
						return
					}
					pid = info
				}
				treasure, err := service.AcquireTreasure(address, tradeHash, pid)
				if !treasure {
					context.JSON(http.StatusOK, ResponseUtil{Msg: err.Error()})
					return
				}
				context.JSON(http.StatusOK, ResponseUtil{Msg: "成功", Status: true})
			}else {
				if err2 != nil {
					context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "hash decode wrong"})
					return
				}

				if len(address) == 0 || len(tradeHash) == 0 {
					context.JSON(http.StatusOK, ResponseUtil{Msg: "no address or tradeHash", Status: false})
					return
				}
				pid := 0
				if len(paddress) > 0 {
					info := sql.ExistsIdByAddress(paddress)
					if info == 0 {
						context.JSON(http.StatusOK, ResponseUtil{Msg: "pid invalid", Status: false})
						return
					}
					pid = info
				}
				treasure, err := service.AcquireTreasure(address, tradeHash, pid)
				if !treasure {
					context.JSON(http.StatusOK, ResponseUtil{Msg: err.Error()})
					return
				}
				context.JSON(http.StatusOK, ResponseUtil{Msg: "success", Status: true})

			}


		})
		//获取价格
		group.GET("/price", func(context *gin.Context) {
			language := context.GetHeader("accept-language")
			price := service.GetPrice()
			if language == "zh-CN" {
				if price == 0 {
					price = 1
				}
				context.JSON(http.StatusOK, ResponseUtil{Msg: "成功", Status: true, Data: gin.H{"price": price}})
			}else {
				if price == 0 {
					price = 1
				}
				context.JSON(http.StatusOK, ResponseUtil{Msg: "success", Status: true, Data: gin.H{"price": price}})
			}
		})
		//发起提现
		group.POST("/cash", func(context *gin.Context) {
			language := context.GetHeader("accept-language")
			address := context.Query("address")
			tradeHashEncode := context.Query("password")
			if len(address)==0&&len(tradeHashEncode)==0 {
				var param Param
				context.BindJSON(&param)
				address = param.Address
				tradeHashEncode = param.Password
			}
			tradeHash, err2 := crypto.BaseDecode(tradeHashEncode)
			if language =="zh-CN"{
				if err2 != nil {
					context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "hash值错误"})
					return
				}
				if len(address) == 0 || len(tradeHash) == 0 {
					context.JSON(http.StatusOK, ResponseUtil{Msg: "地址或者交易hash值无效 ", Status: false})
					return
				}
				out, err := service.CashOut(address, tradeHash)
				if err != nil || !out {
					context.JSON(http.StatusOK, ResponseUtil{Msg: err.Error(), Status: false})
					return
				}
				context.JSON(http.StatusOK, ResponseUtil{Msg: "提现成功", Status: true})

			}else {
				if err2 != nil {
					context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "hash decode wrong"})
					return
				}
				if len(address) == 0 || len(tradeHash) == 0 {
					context.JSON(http.StatusOK, ResponseUtil{Msg: "address or tradeHash invalid ", Status: false})
					return
				}
				out, err := service.CashOut(address, tradeHash)
				if err != nil || !out {
					context.JSON(http.StatusOK, ResponseUtil{Msg: err.Error(), Status: false})
					return
				}
				context.JSON(http.StatusOK, ResponseUtil{Msg: "cash success", Status: true})
			}
	})
		//每日定时任务发起分配
		group.GET("/time", func(context *gin.Context) {
			context.Query("start")
			log.Println("set time ")
			service.Today()
		})
		//参与基金
		//group.POST("/joinFund", func(context *gin.Context) {
		//	address := context.Query("address")
		//	tradeHashEncode := context.Query("password")
		//	tradeHash, err2 := crypto.BaseDecode(tradeHashEncode)
		//	if err2 !=nil{
		//		context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "hash decode wrong"})
		//		return
		//	}
		//	if len(address) ==0 || len(tradeHash) ==0 {
		//		context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "fail" })
		//		return
		//	}
		//	fund, _ := service.JoinFundWithTreasure(address, tradeHash)
		//	if  !fund{
		//		context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "fail"})
		//		return
		//	}
		//	context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "success"})
		//})
		//参与投资
		group.POST("/invest", func(context *gin.Context) {
			language := context.GetHeader("accept-language")
			tradeHashEncode := context.Query("password")
			address := context.Query("address")
			amount := context.Query("amount")
			paddress := context.Query("paddress")
			amount1 :=float32(0)
			if len(tradeHashEncode) == 0 && len(address) == 0 && len(amount) == 0{
				var param Param
				context.BindJSON(&param)
				address = param.Address
				amount1 = param.Amount
				tradeHashEncode = param.Password
				paddress = param.Paddress
			}
			log.Printf("invest  address:%s , amount:%d",address,int(amount1))
			tradeHash, err2 := crypto.BaseDecode(tradeHashEncode)
			if language =="zh-CN"{
				if err2 != nil {
					context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "hash值错误"})
					return
				}
				if len(tradeHash) == 0 || len(address) == 0 {
					context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "参数错误"})
					return
				}
				//amount1, err := strconv.ParseFloat(amount, 32)
				//if err != nil {
				//	context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: ""})
				//	return
				//}
				if amount1 < 50 {
					context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "参与不足50"})
					return
				}
				pid := 0
				if len(paddress) > 0 {
					info := sql.ExistsIdByAddress(paddress)
					if info == 0 {
						context.JSON(http.StatusOK, ResponseUtil{Msg: "推荐人地址无效", Status: false})
						return
					}
					pid = info
				}
				price := service.GetPrice()
				if price == 0 {
					price = 1
				}
				invest, err := service.Invest(address, tradeHash, float32(amount1), price, pid)
				if !invest {
					context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: err.Error()})
					return
				}
				context.JSON(http.StatusOK, ResponseUtil{Status: true, Msg: "成功"})
			}else {
				if err2 != nil {
					context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "Hash decode wrong"})
					return
				}
				if len(tradeHash) == 0 || len(address) == 0 {
					context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "param fail"})
					return
				}
				//amount1, err := strconv.ParseFloat(amount, 32)
				//if err != nil {
				//	context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: ""})
				//	return
				//}
				if amount1 < 50 {
					context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "invest not enough 50"})
					return
				}
				pid := 0
				if len(paddress) > 0 {
					info := sql.ExistsIdByAddress(paddress)
					if info == 0 {
						context.JSON(http.StatusOK, ResponseUtil{Msg: "pid invalid", Status: false})
						return
					}
					pid = info
				}
				price := service.GetPrice()
				if price == 0 {
					price = 1
				}
				invest, err := service.Invest(address, tradeHash, float32(amount1), price, pid)
				if !invest {
					context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: err.Error()})
					return
				}
				context.JSON(http.StatusOK, ResponseUtil{Status: true, Msg: "success"})
			}
		})
		//使用令牌参与投资
		group.POST("/investUseT", func(context *gin.Context) {
			language := context.GetHeader("accept-language")
			address := context.Query("address")
			tradeHashEncode := context.Query("password")
			value := context.Query("value")

			if len(address) == 0 && len(tradeHashEncode) == 0 {
				var param Param
				context.BindJSON(&param)
				address = param.Address
				tradeHashEncode = param.Password
				value = param.Value
			}
			log.Printf("useT address:%s ,anount:%s",address,value)
			tradeHash, err2 := crypto.BaseDecode(tradeHashEncode)
		if language == "zh-CN"{
			if err2 != nil {
				context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "hash值错误"})
				return
			}
			if len(address) == 0 || len(tradeHash) == 0  {
				context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "地址或者交易hash无效"})
				return
			}
			price := service.GetPrice()
			if price == 0 {
				price = 1
			}
			if value == "10" {
				invest, err := service.InvestByTreasure(address, tradeHash, float32(10), price)
				if !invest {
					context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: err.Error()})
					return
				}
				context.JSON(http.StatusOK, ResponseUtil{Status: true, Msg: "成功"})
				return
			}
			if value == "100" {
				fund, _ := service.JoinFundWithTreasure(address, tradeHash)
				if !fund {
					context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "失败"})
					return
				}
				context.JSON(http.StatusOK, ResponseUtil{Status: true, Msg: "成功"})
				return
			} else {
				context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "传入值无效"})
			}
		}else {
			if err2 != nil {
				context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "hash decode wrong"})
				return
			}
			if len(address) == 0 || len(tradeHash) == 0  {
				context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "address or tradeHash wrong"})
				return
			}
			price := service.GetPrice()
			if price == 0 {
				price = 1
			}
			if value == "10" {
				invest, err := service.InvestByTreasure(address, tradeHash, float32(10), price)
				if !invest {
					context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: err.Error()})
					return
				}
				context.JSON(http.StatusOK, ResponseUtil{Status: true, Msg: "success"})
				return
			}
			if value == "100" {
				fund, _ := service.JoinFundWithTreasure(address, tradeHash)
				if !fund {
					context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "fail"})
					return
				}
				context.JSON(http.StatusOK, ResponseUtil{Status: true, Msg: "success"})
				return
			} else {
				context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "value invalid"})
				return
			}
		}
		})
		//使用令牌增加后面的50%收益
		group.POST("/useTreasure", func(context *gin.Context) {
			language := context.GetHeader("accept-language")
			var param Param
			 address := context.PostForm("address")
			if len(address)==0 {
				context.BindJSON(&param)
				address = param.Address
			}
			if language =="zh-CN"{
				if len(address) == 0 {
					context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "钱包地址无效"})
					return
				}
				treasure, err := service.UseTreasure(address)
				if !treasure {
					context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: err.Error()})
					return
				}
				context.JSON(http.StatusOK, ResponseUtil{Status: true, Msg: "成功"})
			}else {
				if len(address) == 0 {
					context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "address invalid"})
					return
				}
				treasure, err := service.UseTreasure(address)
				if !treasure {
					context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: err.Error()})
					return
				}
				context.JSON(http.StatusOK, ResponseUtil{Status: true, Msg: "success"})
			}
		})
		group.GET("/team", func(context *gin.Context) {
			language := context.GetHeader("accept-language")
			address := context.Query("address")
		if language =="zh-CN"{
			if len(address) == 0 {
				context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "钱包地址无效"})
				return
			}

			allNumber, teamNumber := service.GetTeamNumber(address)
			number := teamNumber
			totalU := service.GetTeamTotalU(address)
			context.JSON(http.StatusOK, ResponseUtil{Status: true, Msg: "成功", Data: gin.H{"teamNumber": number, "allNumber": allNumber,"teamPower":totalU}})
			return
		}else {
			if len(address) == 0 {
				context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: "address invalid"})
				return
			}

			allNumber, teamNumber := service.GetTeamNumber(address)
			number := teamNumber
			totalU := service.GetTeamTotalU(address)
			context.JSON(http.StatusOK, ResponseUtil{Status: true, Msg: "success", Data: gin.H{"teamNumber": number, "allNumber": allNumber,"teamPower":totalU}})
			return
		}
		})
		group.GET("/detail", func(context *gin.Context) {
			language := context.GetHeader("accept-language")
			address := context.Query("address")
			if language =="zh-CN"{
				if len(address) == 0 {
					context.JSON(http.StatusNotFound, ResponseUtil{Msg: "钱包地址无效", Status: false})
					return
				}
				price := service.GetPrice()
				if price == 0 {
					price = 1
				}
				info := service.GetUserInfo(address, price)
				context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "成功",Data: gin.H{"IsHasTreasure": info.Treasure, "item": gin.H{"fund": info.Fund, "power": info.TotalU}}})
			} else {
				if len(address) == 0 {
					context.JSON(http.StatusNotFound, ResponseUtil{Msg: "address invalid", Status: false})
					return
				}
				price := service.GetPrice()
				if price == 0 {
					price = 1
				}
				info := service.GetUserInfo(address, price)
				context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "success",Data: gin.H{"IsHasTreasure": info.Treasure, "item": gin.H{"fund": info.Fund, "power": info.TotalU}}})
			}
		})
	}
}

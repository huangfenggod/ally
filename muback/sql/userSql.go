package sql

import (
	"errors"
)

type BindAddress struct {
	Address string `db:"address"`
	Bind string `db:"bind"`
	TotalU float32 `db:"total_u"`
	Fund float32 `db:"fund"`
	BalanceU float32 `db:"baclanceU"`
	Allowance int64 `db:"allowance"`
}
func GetUserListByAddress(address string,pageSize int,pageNumber int)(int,[]BindAddress,error)  {
	var u []BindAddress
	var u1 User
	var c count
	DB.Raw("select uid from user where address =?",address).Scan(&u1)
	if u1.Uid==0 {
		return 0,u,errors.New("dont have this address")
	}

	DB.Raw("with recursive cte as (select uid,address,pid,total_u,fund from user where pid = ? union all select u.uid,u.address,u.pid,u.total_u,u.fund from cte c inner join user u on c.uid=u.pid) select count(*) cou  from cte",u1.Uid).Scan(&c)
	//DB.Raw("with recursive cte as (select uid,address,pid,total_u,fund from user where pid = ? union all select u.uid,u.address,u.pid,u.total_u,u.fund from cte c inner join user u on c.uid=u.pid)select u1.address,u1.total_u,u1.fund,u2.bind from (select uid,pid,address,total_u,fund from cte)as u1  inner join (select uid,address bind from user)as u2 on u1.pid=u2.uid ",u1.Uid).Scan(&u)
	DB.Raw("with recursive cte as (select uid,address,pid,total_u,fund from user where pid = ? union all select u.uid,u.address,u.pid,u.total_u,u.fund from cte c inner join user u on c.uid=u.pid)select u1.address,u1.total_u,u1.fund,u2.bind from (select uid,pid,address,total_u,fund from cte)as u1  inner join (select uid,address bind from user)as u2 on u1.pid=u2.uid limit ?,? ",u1.Uid,(pageNumber-1)*pageSize,pageSize).Scan(&u)
	return c.Cou,u,nil
}

func GetAllUserList(pageSize int,pageNumber int) (int,[]BindAddress) {
	var u []BindAddress
	var c count
	DB.Raw("select count(*) cou from user where authority=0").Scan(&c)
	DB.Raw("select t1.address address,t1.total_u,t1.fund ,t2.address bind from (select uid,Address,pid ,total_u,fund from user where authority=0) t1 left join (select uid ,address from user)t2 on t1.pid=t2.uid limit ?,?",(pageNumber-1)*pageSize,pageSize).Scan(&u)
	return c.Cou,u
}

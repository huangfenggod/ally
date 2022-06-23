package sql

import "errors"

func LoginUserPasswd(username string)*Admin{
	var ad Admin
	DB.Raw("select * from admin where username =? and disable=1",username).Scan(&ad)
	return &ad
}

type exists struct {
	Ex int `db:"ex"`
}
func CreateUser(username string,password string,bindAddress string)(bool,error) {
	var exi exists
	DB.Raw("select 1 ex from admin where username =? limit 1 ", username).Scan(&exi)
	if exi.Ex > 0 {
		return false, errors.New("username has exists")
	}

	var exi1 exists
	DB.Raw("select 1 ex from user where address =? limit 1",bindAddress).Scan(&exi1)
	if exi1.Ex==0 {
		return false,errors.New("bindAddress not exists")
	}

	exec := DB.Exec("insert admin(username,password,bind,role) values(?,?,?,'ordinary')", username, password, bindAddress)
	if exec.Error!=nil {
		return false,exec.Error
	}
	return true,nil
}
func DeleteUser(username string )  {
	var exi exists
	DB.Raw("select 1 ex from admin where username =? limit 1 ", username).Scan(&exi)
	if exi.Ex == 0 {
		return
	}
	//DB.Exec("delete from user where username=?",username)
	DB.Exec("update admin set disable=0 where role = 'ordinary' and username=?",username)
}

type count struct {
	Cou int `db:"cou"`
}
func GetAdminNumber() int {
	var co count
	DB.Raw("select count(*) cou from admin where role = 'ordinary' and disable=1 ").Scan(&co)
	return co.Cou
}


func GetUserList(pageSize int,pageNumber int) []Admin{
	var admin []Admin
	DB.Raw("select username,password,bind,role from admin where role = 'ordinary' and disable=1 limit ?,?",(pageNumber-1)*pageSize,pageSize).Scan(&admin)
	return admin
}

func GetBindAddress(username string)string {
	var ad Admin
	DB.Raw("select bind from admin where username =? ",username).Scan(&ad)
	return ad.Bind
}

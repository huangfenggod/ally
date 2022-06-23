package crypto

import "errors"

func BaseDecode(str string)(string,error ) {
	var  s string
	i := len(str)
	if i <10 {
		return s,errors.New("hash wrong")
	}
	s = str[:10]+str[20:]
	return s,nil
}


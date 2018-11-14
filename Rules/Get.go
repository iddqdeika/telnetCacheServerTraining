package Rules

import (
	"errors"
	"net"
)

const GetKey = "get"

type GetRule struct{
	Cache Cache
}

func (r *GetRule) Meets(message string) bool{
	phrases := SplitIntoPhrases(message)
	if (phrases[0]==GetKey && len(phrases)==2){
		return true
	}
	return false
}

func (r *GetRule) Process(message string, conn net.Conn) error{
	if r.Cache!=nil {
		phrases := SplitIntoPhrases(message)
		str, err := r.Cache.Get(phrases[1])
		if err!=nil{
			conn.Write([]byte("0"))
			return err
		}
		conn.Write([]byte("1\r\n" + str))
		return err
	}else{
		return errors.New("no cache found for operation " + GetKey)
	}
}

func (r *GetRule) SetCache(cache Cache){
	r.Cache = cache
}

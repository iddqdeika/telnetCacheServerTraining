package Rules

import (
	"errors"
	"net"
)

const DelKey = "del"

type DelRule struct{
	Cache Cache
}

func (r *DelRule) Meets(message string) bool{
	phrases := SplitIntoPhrases(message)
	if (phrases[0]==DelKey && len(phrases)==2){
		return true
	}
	return false
}

func (r *DelRule) Process(message string, conn net.Conn) error{
	if r.Cache!=nil {
		phrases := SplitIntoPhrases(message)
		err := r.Cache.Delete(phrases[1])
		if err!=nil{
			conn.Write([]byte("0"))
			return err
		}
		_, err = conn.Write([]byte("1"))
		return err
	}else{
		return errors.New("no cache found for operation " + DelKey)
	}
}

func (r *DelRule) SetCache(cache Cache){
	r.Cache = cache
}

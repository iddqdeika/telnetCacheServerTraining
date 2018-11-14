package Rules

import (
	"errors"
	"net"
)

const SetKey =  "set"

type SetRule struct{
	Cache Cache
}

func (r *SetRule) Meets(message string) bool{
	phrases := SplitIntoPhrases(message)
	if (len(phrases)>0&&phrases[0]==SetKey && len(phrases)==3){
		return true
	}
	return false
}

func (r *SetRule) Process(message string, conn net.Conn) error{
	if r.Cache!=nil {
		phrases := SplitIntoPhrases(message)
		key := phrases[1]
		value := phrases[2]
		err := r.Cache.Put(key, value)
		if err!=nil{
			return err
		}
		_, err = conn.Write([]byte("OK"))
		return err
	}else{
		return errors.New("no cache found for operation " + SetKey)
	}
}

func (r *SetRule) SetCache(cache Cache){
	r.Cache = cache
}
package Rules

import (
	"bytes"
	"net"
)

type Cache interface {
	Contains(key string) bool
	Put(key string,value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type Rule interface {
	Meets(message string) bool
	Process(message string, conn net.Conn) error
	SetCache(cache Cache)
}

func SplitIntoPhrases(message string) []string{
	res := make([]string,0)
	incomma := false
	buf := ""

	for _,v := range bytes.Runes([]byte(message)){
		char := string(v)
		if char == "\b"{
			if len(buf)>0{
				buf = buf[:len(buf)-1]
				continue
			}
		}else{
			buf = buf + char
		}
	}
	temp := buf
	buf = ""
	for _,v := range bytes.Runes([]byte(temp)){
		char := string(v)
		println(char)

		if incomma{
			switch char {
			case "\"":
				if len(buf)>0 { res = append(res,buf)}
				buf = ""
				incomma = false
				break
			default:
				buf = buf + char
			}
		}else{
			switch char {
			case " ":
				if len(buf)>0 { res = append(res,buf)}
				buf = ""
				break
			case "\"":
				incomma = true
				if len(buf)>0 { res = append(res,buf)}
				buf = ""
				break
			default:
				buf = buf + char
			}
		}
	}
	if len(buf)>0 { res = append(res,buf)}
	return res
}

func GetRuleSet(cache Cache) []Rule{
	res := make([]Rule,0)
	res = append(res, &SetRule{Cache:cache})
	res = append(res, &GetRule{Cache:cache})
	res = append(res, &DelRule{Cache:cache})
	res = append(res,&ExistsRule{Cache:cache})
	return res
}
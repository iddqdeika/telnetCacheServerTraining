package Rules

import (
	"errors"
	"net"
)

const ExistsKey = "exists"

type ExistsRule struct {
	Cache Cache
}

func (r *ExistsRule) Meets(message string) bool {
	phrases := SplitIntoPhrases(message)
	if phrases[0] == ExistsKey && len(phrases) == 2 {
		return true
	}
	return false
}

func (r *ExistsRule) Process(message string, conn net.Conn) error {
	if r.Cache != nil {
		phrases := SplitIntoPhrases(message)
		exists := r.Cache.Contains(phrases[1])
		if exists {
			conn.Write([]byte("1"))
		} else {
			conn.Write([]byte("0"))
		}
		return nil
	} else {
		return errors.New("no cache found for operation " + ExistsKey)
	}
}

func (r *ExistsRule) SetCache(cache Cache) {
	r.Cache = cache
}

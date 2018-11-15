package Service

import (
	"bufio"
	"net"

	"homework5/Cache"
	"homework5/Rules"
)

type Rule interface {
	Meets(message string) bool
	Process(message string, conn net.Conn) error
	SetCache(cache Rules.Cache)
}

type Service struct {
	Cache   Rules.Cache
	RuleSet []Rule
}

func (s *Service) Accept(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("\u001B[2J"))
	for {
		var bts []byte
		scanner := bufio.NewScanner(conn)
		scanner.Scan()

		bts = scanner.Bytes()
		part := string(bts)
		s.Proceed(part, conn)
		if part == "quit" {
			break
		}
	}
}

func (s *Service) Proceed(text string, conn net.Conn) {
	defer conn.Write([]byte("\r\n"))
	conn.Write([]byte("\u001B[2J"))
	for _, rule := range s.RuleSet {
		if rule.Meets(text) {
			rule.Process(text, conn)
			return
		}
	}
	conn.Write([]byte("cant find rule for commant \"" + text + "\""))
}

func NewService() *Service {
	cache := Cache.GetNewRamCache()
	return &(Service{Cache: &cache, RuleSet: Rules.GetRuleSet(&cache)})
}

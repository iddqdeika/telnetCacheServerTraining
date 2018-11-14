package main

import (
	"bufio"
	"homework5/Rules"
	"net"
)

func GetNewService() Service{
	cache := GetNewRamCache()
	return Service{Cache: &cache,RuleSet:Rules.GetRuleSet(&cache)}
}

type Service struct{
	Cache Rules.Cache
	RuleSet []Rules.Rule
}

func (s *Service) Accept(conn net.Conn){
	defer conn.Close()
	conn.Write([]byte("\u001B[2J"))
	for {
		var bts []byte
		scanner := bufio.NewScanner(conn)
		scanner.Scan()

		bts = scanner.Bytes()
		part := string(bts)
		s.Proceed(part,conn)
		if part=="quit"{
			break
		}
	}
}

func (s *Service) Proceed(text string, conn net.Conn)  {
	defer conn.Write([]byte("\r\n"))
	conn.Write([]byte("\u001B[2J"))
	for _,rule := range s.RuleSet{
		if rule.Meets(text){
			rule.Process(text,conn)
			return
		}
	}
	conn.Write([]byte("cant find rule for commant \"" + text + "\""))
}


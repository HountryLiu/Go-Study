package main

import (
	"errors"
	"fmt"
)

type Server struct {
	curIndex int
	res      []string
}

func (s *Server) Add(addr string) (err error) {
	if addr == "" {
		err = errors.New("addr is null")
	}
	s.res = append(s.res, addr)
	return
}

func (s *Server) Get() (res string, err error) {
	if len(s.res) == 0 {
		err = errors.New("Server is null")
		return
	}
	//随机负载
	// rand.Seed(time.Now().UnixNano())
	// s.curIndex = rand.Intn(len(s.res))
	// res = s.res[s.curIndex]
	//轮询负载
	res = s.res[s.curIndex]
	s.curIndex = (s.curIndex + 1) % len(s.res)
	return
}
func main() {
	s := new(Server)

	s.Add("s1")
	s.Add("s2")
	s.Add("s3")
	s.Add("s4")

	fmt.Println(s.Get())
	fmt.Println(s.Get())
	fmt.Println(s.Get())
	fmt.Println(s.Get())
	fmt.Println(s.Get())
	fmt.Println(s.Get())
	fmt.Println(s.Get())
	fmt.Println(s.Get())
}

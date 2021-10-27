package main

import "zinx/znet"

func main() {
	s := znet.NewServer("[Zinx v0.1]")
	s.Serve()
}



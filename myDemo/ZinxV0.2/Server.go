package main

import "zinx/znet"

func main() {
	s := znet.NewServer("[Zinx v0.2]")
	s.Serve()
}



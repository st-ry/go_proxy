package main

import (
	"fmt"
	"go_proxy/dist"
	"go_proxy/proxy"
)

func main() {
	dist.ServeDist()
	fmt.Println("dist")
	proxy.ServeProxy()
	fmt.Println("hi")
}
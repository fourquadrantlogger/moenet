package main

import (
	"fmt"
	"github.com/timeloveboy/moenet"
)

func main() {
	c := moenet.MoeClient

	fmt.Println(c.Get("http://localhost:9000/index"))
	//fmt.Println(http.DefaultClient.Get("http://localhost:9000/index"))
	//c.Get("http://localhost:9000/relocation")
}

package main

import (
	"fmt"
	"github.com/timeloveboy/moenet"
)

func main() {
	c := moenet.NewClient()

	fmt.Println(c.Do("Get", "http://localhost:9000/index", nil))
	//fmt.Println(http.DefaultClient.Get("http://localhost:9000/index"))
	//c.Get("http://localhost:9000/relocation")
}

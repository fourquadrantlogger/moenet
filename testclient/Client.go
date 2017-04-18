package main

import (
	"fmt"
	"github.com/timeloveboy/moenet"
)

func main() {
	c := moenet.NewClient()

	fmt.Println(c.Do(moenet.MakeRequest("GET").Url("http://localhost:9000/index")))
	//fmt.Println(http.DefaultClient.Get("http://localhost:9000/index"))
	//c.Get("http://localhost:9000/relocation")
}

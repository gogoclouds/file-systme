package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// nohup ./file-system -port 83 -path /home/xxx &

// 启动参数
var port = flag.Int("port", 80, "app prot")
var dir = flag.String("path", ".", "root dir")

func main() {
	flag.Parse()
	log.Printf("port=%d, dir='%s'", *port, *dir)

	addr := fmt.Sprintf(":%d", *port)
	if err := http.ListenAndServe(addr, http.FileServer(http.Dir(*dir))); err != nil {
		log.Fatalln("http.ListenAndServe: ", err)
	}
}

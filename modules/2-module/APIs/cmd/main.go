package main

import (
	"log"

	"github.com/rafaelcamelo31/graduate-go-course/2-module/APIs/configs"
)

func main() {
	conf, _ := configs.LoadConfig(".")
	log.Println(conf.Driver)
}

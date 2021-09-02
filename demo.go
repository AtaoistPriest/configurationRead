package main

import (
	"log"
	"me/demo/config"
)

func main(){
	configs := config.InitConfigs()
	log.Print(configs.Cfgs.Ip)
	log.Print(configs.Cfgs.MAC)
	log.Print(configs.Cfgs.MTU)
	log.Print(configs.Cfgs.Judge)
	log.Print(configs.Cfgs.Time)
}

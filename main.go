package main

import (
	"flag"
	"github.com/robfig/cron/v3"
	"log"
	"sign_http/logic/cc"
	"sign_http/logic/kanxue"
	"sign_http/logic/qile"
)

var clockInConf = flag.String("clockin", "0 12 * * *", "cron 语法的定时签到")

func main() {
	AccessCollection()

	if *clockInConf != "" {

		scheduler := cron.New()
		_, err := scheduler.AddFunc(*clockInConf, func() {
			AccessCollection()

		})
		if err != nil {
			log.Fatal(err)
		}
		scheduler.Run()
	}
}

func AccessCollection() {
	cc.CCSign()
	kanxue.KanXueSign()
	qile.QiLeSign()
}

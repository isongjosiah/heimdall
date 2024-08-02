package logic

import (
	"github.com/robfig/cron/v3"
	"log"
)

func InitCron(l *Logic) {

	client := cron.New()
	if _, err := client.AddFunc("@every 1h", l.Monitor.RetrieveCommit); err != nil {
		log.Fatal("Unable to start repository monitoring ", err)
	}
	client.Start()

}

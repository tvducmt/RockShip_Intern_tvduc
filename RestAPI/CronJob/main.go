package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	c.AddFunc("1 * * * * *", func() {
		fmt.Println(time.Now()) //in thoi gian hien tai sau mot phuc
	})
	fmt.Println(" Start ...")
	c.Run()

	//c.Stop()
}

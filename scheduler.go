package main

import (
	"fmt"
	"github.com/carlescere/scheduler"
	 _ "net/http/pprof"
)

type Curtain struct {
	startDate string
	endDate string
	FadeJob *scheduler.Job
}

func execute(start string,end string,interval int,job func()) {
	curtain := Curtain{startDate: "test",endDate:"test"}
	fmt.Println(curtain)

	// Run every 2 seconds but not now.
	curtain.FadeJob, _ = scheduler.Every(interval).Seconds().NotImmediately().Run(job)
	curtain.FadeJob, _ = scheduler.Every(30).Seconds().NotImmediately().Run(curtain.quit)
	scheduler.Every().Day().At(start).Run(job)
	scheduler.Every().Day().At(end).Run(curtain.quit)

}

func (c Curtain)quit(){
	c.FadeJob.Quit <- true
}
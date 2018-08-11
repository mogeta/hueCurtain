package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/carlescere/scheduler"
)

func main() {
	job := func() {
		t := time.Now()
		fmt.Println("Time's up! @", t.UTC())
	}


	// Run every 2 seconds but not now.
	scheduler.Every(2).Seconds().NotImmediately().Run(fade)

	//j.Quit <- true

	scheduler.Every().Monday().At("19:00").Run(job)
	scheduler.Every().Monday().At("23:30").Run(job)
	// Run now and every X.
	//scheduler.Every(5).Minutes().Run(job)
	//scheduler.Every().Day().Run(job)
	//scheduler.Every().Monday().At("08:30").Run(job)


	// Keep the program from not exiting.
	runtime.Goexit()
}

func fade(){
	t := time.Now()
	fmt.Println("Time's up! @", t.UTC())
}
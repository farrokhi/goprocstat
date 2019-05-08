package main

import (
	"flag"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"time"
)

func main() {

	var currUser, prevUser, currSystem, prevSystem, cpuUsage float64

	flagPID := flag.Int("p", 1, "PID for process to monitor")
	flagNUM := flag.Int("n", -1, "Stop after n times of displaying stat")
	flag.Parse()

	p, _ := process.NewProcess(int32(*flagPID))
	name, _ := p.Name()
	cnt := 0

	v, _ := p.Times()
	currUser = v.User
	currSystem = v.System

	for !(*flagNUM >= 0 && cnt >= *flagNUM) {

		v, err := p.Times()

		if err == nil {
			// keep old counters, needed to calculated variance
			prevUser = currUser
			prevSystem = currSystem

			currUser = v.User
			currSystem = v.System

			cpuUsage, _ = p.CPUPercent()

			fmt.Printf("%d  %s  cpu: %5.2f%%   user: %5.2f (%+5.2f)   system: %5.2f (%+5.2f)   iowait: %5.2f   irq: %5.2f   softirq: %5.2f\n",
				time.Now().Unix(), name, cpuUsage, currUser, currUser-prevUser, currSystem, currSystem-prevSystem, v.Iowait, v.Irq, v.Softirq)
		} else {
			fmt.Println(err)
		}
		time.Sleep(time.Second)
		cnt++
	}
}

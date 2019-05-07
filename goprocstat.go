package main

import (
	"flag"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"os"
	"time"
)

func main() {

	flagPID := flag.Int("p", 1, "PID for process to monitor")
	flagNUM := flag.Int("n", -1, "Stop after n times of displaying stat")
	flag.Parse()

	p, _ := process.NewProcess(int32(*flagPID))
	v, err := p.Times()
	name, _ := p.Name()
	cnt := 0

	for {
		if *flagNUM >= 0 && cnt >= *flagNUM {
			os.Exit(0)
		}

		if err == nil {
			cpuUsage, _ := p.CPUPercent()
			fmt.Printf("%d  %s  cpu: %3.2f%%   user: %-3.2f   system: %0.2f   iowait: %0.2f   irq: %0.2f   softirq: %0.2f\n",
				time.Now().Unix(), name, cpuUsage, v.User, v.System, v.Iowait, v.Irq, v.Softirq)
		} else {
			fmt.Println(err)
		}
		time.Sleep(time.Second)
		cnt++
	}
}

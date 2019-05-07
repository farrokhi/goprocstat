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
	cnt:=0

	for {
		if *flagNUM >= 0  &&  cnt >= *flagNUM {
				os.Exit(0)
			}

		if err == nil {
			fmt.Printf("%d  pid: %d   user: %f\t system: %f\n", time.Now().Unix(), *flagPID, v.User, v.System)
		} else {
			fmt.Println(err)
		}
		time.Sleep(time.Second)
		cnt++
	}
}

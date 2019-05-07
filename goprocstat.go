package main

import (
	"flag"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"time"
)

func main() {

	flagPID := flag.Int("p", 1, "Gimme your PID")
	flag.Parse()

	p, _ := process.NewProcess(int32(*flagPID))
	v, err := p.Times()

	for {
		if err == nil {
			fmt.Printf("%d  pid: %d   user: %f\t system: %f\n", time.Now().Unix(), *flagPID, v.User, v.System)
		} else {
			fmt.Println(err)
		}
		time.Sleep(time.Second)
	}

}

package main

import "time"

func sleep(s time.Duration) {
	select {
	case <-time.After(s):
		return
	}
}

func sleep2(s time.Duration) {
	timer := time.NewTimer(s)
	<-timer.C
}

func main() {
	t := 2
	sleep(time.Duration(t) * time.Second)
	sleep2(time.Duration(t) * time.Second)
}

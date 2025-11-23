package main

import (
	"fmt"
	"time"
)

func MySleep(d time.Duration) {
	timer := time.NewTimer(d)
	<-timer.C
}

func main() {
	fmt.Println("start", time.Now())
	MySleep(2 * time.Second)
	fmt.Println("end  ", time.Now())
}
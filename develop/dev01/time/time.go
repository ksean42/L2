package main

import (
	"fmt"

	"github.com/beevik/ntp"
)

func main() {
	time, _ := ntp.Time("0.beevik-ntp.pool.ntp.org")
	fmt.Println(time)
}

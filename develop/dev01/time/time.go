package time

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func Time() {
	if time, err := ntp.Time("0.beevik-ntp.pool.ntp.org"); err != nil {
		fmt.Println(time)
	} else {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}

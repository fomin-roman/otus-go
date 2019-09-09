package hw_1

import (
	"fmt"
	"time"
	"github.com/beevik/ntp"
)

func printCurrentTime() error {

	currentTime := time.Now()

	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		return fmt.Errorf("ERROR : %s", err.Error())
	}

	fmt.Println("Current time : " , currentTime)
	fmt.Println("Exact time : " , exactTime)

	return nil
}

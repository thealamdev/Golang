package times

import (
	"fmt"
	"time"
)

func GetTime() {
	currentTime := time.Now()
	fmt.Println(currentTime.Format("01 02 2006 15:04:05 Monday"))
}

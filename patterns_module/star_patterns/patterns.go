package star_patterns

import (
	"fmt"
	"time"
)

func PrintPattern(rows int) {
	timeDelay := 2
	for j := rows; j >= 1; j-- {
		for i := j; i >= 1; i-- {
			fmt.Print("*")
			time.Sleep(time.Duration(timeDelay) * time.Millisecond)
		}
		fmt.Println("")
	}

	for j := 1; j < rows; j++ {
		for i := 0; i < j; i++ {
			fmt.Print("*")
			time.Sleep(time.Duration(timeDelay) * time.Millisecond)
		}
		fmt.Println("")
	}
}

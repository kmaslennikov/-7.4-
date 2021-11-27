package main

import (
	"fmt"
)

const START = 100000
const END = 999999

func main() {
	max := 0
	ticketNum := START
	for i := START; i <= END; i++ {
		sum1 := i/100000 + (i/10000)%10 + (i/1000)%10
		sum2 := (i/100)%10 + (i/10)%10 + i%10

		if sum1 == sum2 {
			fmt.Println(i)
			fmt.Println(max)
			fmt.Println(ticketNum)
			fmt.Println()
			if max < i-ticketNum {
				max = i - ticketNum
				ticketNum = i
			}
		}
	}
	fmt.Println("Минимальное количество билетов, которое нужно купить, чтобы среди них оказался счастливый:", max)
}

package main

import "fmt"

func corpFlightBookings(bookings [][]int, n int) []int {
	seats := make([]int, n)
	for _, b := range bookings {
		seats[b[0]-1] += b[2]
		if b[1] < n {
			seats[b[1]] -= b[2]
		}
	}
	for i := 1; i < n; i++ {
		seats[i] += seats[i-1]
	}
	return seats
}

func main() {
	bookings := [][]int{
		{1, 2, 10},
		{2, 3, 20},
		{2, 5, 25},
	}
	n := 5
	fmt.Println(corpFlightBookings(bookings, n))
}

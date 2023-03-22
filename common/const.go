package common

import "fmt"

const (
	DBTypeRestaurant = 1
	DBTypeUser       = 2
)

func AppRecover() {
	if err := recover(); err != nil {
		fmt.Printf("error recovered: %s", err)
	}
}

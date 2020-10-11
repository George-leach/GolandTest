package main

import "time"
import "fmt"

func main() {
	go heavy()
	go superHeavy()
	fmt.Println("go testing")
	select {}
}
func superHeavy() {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("super heavy")
	}
}
func heavy() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Hello")
	}

}

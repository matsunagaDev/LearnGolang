package main

import (
	"fmt"
	"os/user"
	"time"
)

/*
func init() {
	fmt.Println("Init!")
}
*/

func bazz() {
	fmt.Println("Bazz")
}

func main() {
	// bazz()
	fmt.Println("Hello, World!", "TEST TEST", time.Now())
	fmt.Println(user.Current())
}
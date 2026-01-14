package main

import (
	"fmt"
	"io"
	"log"
	"os"
)
func incrementGenerator() (func() int) {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
 	}
}

func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

func by2(num int) string {
	if num % 2 == 0 {
		return "ok" 
	} else {
		return "no"
	}
}

func getOsName() string{
	return "aiueo"
}

func foo2() {
	defer fmt.Println("world foo")

	fmt.Println("hello foo")
}

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR | os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log. Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}


func main() {
	LoggingSettings("test.log")
	_, err := os.Open("fatadaf;asf")
	if err != nil {
		log.Fatalln("Exit", err)
	}
	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	log.Fatalf("%T %v", "test", "test")
	log.Fatalln("error!")

	fmt.Println("ok!")
}

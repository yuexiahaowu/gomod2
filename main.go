package main

import (
	"fmt"
	_ "github.com/apex/log"
	"time"
)

func main() {
	time.Sleep(100)
	fmt.Println("  gomod 2   main ")
}

func init(){
	time.Sleep(1000000)
	fmt.Println(" gomod2   init ")
}


package main

import (
	"fmt"
	"lesson8/config"
)

// this is a comment

func main() {

	// confEnv := config.GetEnv()
	// fmt.Println(confEnv)

	// confFlags := config.GetFlags()
	// fmt.Println(confFlags)

	conf := config.GetConfig()

	fmt.Println(conf)
}

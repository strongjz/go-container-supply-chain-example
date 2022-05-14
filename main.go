package main

import (
	"fmt"
	"github.com/strongjz/go-container-supply-chain-example/app"
)

func main() {

	app := app.New()
	fmt.Println("Starting App")
	app.Engine()
	app.Start()
}

package main

import (
	"fmt"
)

func main() {

	app := &App{}
	fmt.Println("Starting App")
	app.Engine()
	if err := app.Start(); err != nil {
		fmt.Println("Error Starting server: ", err)
	}
}

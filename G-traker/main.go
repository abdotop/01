package main

import (
	"G-tracker/internal/handelers"
	"fmt"
	"net/http"
)

func main() {
	handelers.Hand()
	fmt.Printf("Your server has successfully started on port %v \n\t(http://localhost%v)\n", Port[1:], Port)
	if err := http.ListenAndServe(Port, nil); err != nil {
		panic(err)
	}
}

const Port string = ":8080"

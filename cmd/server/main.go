package main

import "fmt"

// Run - responsible for instantiation and startup of the app
func Run() error {
	fmt.Println("Starting up app")
	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}

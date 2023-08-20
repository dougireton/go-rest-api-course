package main

import "fmt"

// Run starts the API app, returning an error if start up isn't successful.
func Run() error {
	fmt.Println("Starting app")
	return nil
}

func main() {
	fmt.Println("Go REST API Course")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}

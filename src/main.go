package main

import "fmt"

func main() {

	releaseState := Version{releaseNumber: "0.0.4", state: "demo"}
	fmt.Println("################################################")
	fmt.Println("Helm Migrate Plugin version is :%s state is : ", releaseState.releaseNumber, releaseState.state)
	fmt.Println("################################################\n")

	Execute()

}

package main

import "fmt"

func main() {

	releaseState := Version{releaseNumber: "0.0.6", state: "demo"}
	fmt.Println("################################################")
	fmt.Println("Helm Migrate Plugin version:: ", releaseState.releaseNumber, "state is:", releaseState.state)
	fmt.Println("################################################\n")

	Execute()

}

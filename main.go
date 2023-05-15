package main

import (
	"fmt"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
	uuid "github.com/google/uuid"
	"rsc.io/quote/v3"
)

func main() {
	say, err := cowsay.Say(
		quote.HelloV3(),
		cowsay.Type("default"),
		cowsay.BallonWidth(60),
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(uuid.New().Domain())
	fmt.Println(say)
}

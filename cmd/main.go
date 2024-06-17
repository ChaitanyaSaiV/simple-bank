package main

import (
	"fmt"

	"github.com/ChaitanyaSaiV/simple-bank/router"
)

func main() {
	fmt.Println("Hello World!!")

	router.InitRouter()
	router.Start(":8080")

}

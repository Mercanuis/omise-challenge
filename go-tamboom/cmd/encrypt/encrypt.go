package main

import (
	"os"

	"go-tamboom/service"
)

func main() {
	s := service.NewChargerService(os.Args[1])
	s.MakeDecipherText()
}

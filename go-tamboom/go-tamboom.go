//Package main contains the main entry point
//This file is meant to be run with the following command
//
//    ./go-tamboom <file_name>
//where <file_name> is a 128 encrypted file
//This process is treated like a batch job: errors are logged, but the application should not crash, or if it does, can be re-ran
package main

import (
	"os"

	"go-tamboom/service"
)

func main() {
	s := service.NewChargerService(os.Args[1])
	s.ProcessDonations()
}

package main

import (
	"flag"
	"fmt"

	"github.com/JanigaAtWork/NifiHealthCheck/GetStatus"
)

func main() {

	var cert string
	var cacert string

	flag.StringVar(&cert, "cert", "", "path to your cert")

	flag.StringVar(&cacert, "cacert", "", "path to your ca cert")

	flag.Parse()

	fmt.Println(GetStatus.GetStatus(cacert, cert))
}

package main

import (
	"github.com/diazharizky/rest-otp-generator/cmd"
	"github.com/diazharizky/rest-otp-generator/configs"
)

func main() {
	configs.LoadConfig()
	cmd.Execute()
}

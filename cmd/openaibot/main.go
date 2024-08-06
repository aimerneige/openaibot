package main

import "github.com/aimerneige/openaibot/internal/bootstrap"

// Variables are overrides during build
// See the Makefile for more information.
var (
	ConfPath   string = "./config/application.yaml"
	DebugLevel string = "debug"
)

func main() {
	bootstrap.InitLog(DebugLevel)
	bootstrap.InitConfig(ConfPath)
	bootstrap.StartBot()
}

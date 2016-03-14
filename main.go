package main

import (
	"os"

	"github.com/go-martini/martini"
)

func getInfo(params martini.Params) (int, []byte) {
	services := os.Getenv("VCAP_SERVICES")
	return 200, []byte(services)
}

func main() {
	m := martini.Classic()

	m.Get("/info", getInfo)
	m.Run()
}

package main

import (
	"gopkg.in/macaron.v1"
)

func main() {

	m := macaron.Classic()
	initRoutes(m)
	m.Run()
}

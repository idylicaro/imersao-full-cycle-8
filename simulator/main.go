package main

import (
	"fmt"

	route2 "github.com/idylicaro/imersao-full-cycle-8/simulator/application/route"
)

func main() {
	route := route2.Route{
		ID:       "1",
		ClientID: "1",
	}
	route.LoadPositions()
	stringJson, _ := route.ExportJsonPositions()
	fmt.Println(stringJson[0])
}

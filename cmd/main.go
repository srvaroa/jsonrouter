package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/srvaroa/jsonrouter/pkg/router"
)

func main() {

	configFile := flag.String("routes", "", "The path of the file containing the routes definition")
	flag.Parse()

	configData, err := ioutil.ReadFile(*configFile)
	if err != nil {
		fmt.Printf("Could not read route configuration: %s\n", err)
		os.Exit(1)
	}

	var routes *router.RoutingTable
	routes, err = router.NewRoutingTable(&configData)
	if err != nil {
		os.Exit(1)
	}

	var data []byte
	data, err = ioutil.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(1)
	}

	json_raw := string(data)
	res, err := routes.FindMatches(&json_raw)
	if err != nil {
		os.Exit(1)
	}

	for endpoint, _ := range res {
		fmt.Println(endpoint)
	}
}

package main

import (
	"fmt"

	"github.com/jackpal/gateway"
)

func main() {
	// Get the IP address of the default gateway
	defaultGateway, err := gateway.DiscoverGateway()
	if err != nil {
		fmt.Println("Could not get default gateway: ", err)
	}

	fmt.Println("Aquired default gateway: ", defaultGateway)

	// I know there should be a better way to do this but, this is mostly a bodge
	err = callPingCommand(defaultGateway)

	if err != nil {
		fmt.Println(err)
	}
}

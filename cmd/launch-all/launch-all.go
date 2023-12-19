package main

import (
	"fmt"

	"gitlab.utc.fr/mennynat/ia04-project/serveur"
)

func main() {
	server := serveur.NewRestServerAgent("localhost:8080")
	server.Start()
	fmt.Scanln()
}

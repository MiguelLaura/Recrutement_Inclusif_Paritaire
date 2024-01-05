/*
Crée le serveur web qui gère toutes les simulations.

Utilisation :

	launch-all [flags]

Les flags peuvent être :

	-h, --host nomHôte
		Indique le nom de l'hôte.
		Défaut : localhost

	-p, --port numeroPort
		Indique le port du serveur.
		Défaut : 8080

	-s
		Lance le serveur en mode silencieux
*/
package main

import (
	"flag"
	"fmt"
	"io"
	"log"

	"gitlab.utc.fr/mennynat/ia04-project/serveur"
)

func main() {

	// -------
	// Traitement des flags

	var host string
	var port int
	var silent bool

	flag.StringVar(&host, "host", "localhost", "Hôte du serveur")
	flag.StringVar(&host, "h", "localhost", "Hôte du serveur (raccourci)")

	flag.IntVar(&port, "port", 8080, "Port du serveur")
	flag.IntVar(&port, "p", 8080, "Port du serveur (raccourci)")

	flag.BoolVar(&silent, "s", false, "Le serveur est en mode silencieux ou non")

	flag.Parse()

	if port < 0 {
		log.Fatalf("Le numéro de port ne peut être négatif (donné %d)", port)
	}

	// ------
	// Execution du script

	if silent {
		log.SetFlags(0)
		log.SetOutput(io.Discard)
	}

	server := serveur.NewRestServerAgent(fmt.Sprintf("%s:%d", host, port))
	server.Start()

	fmt.Scanln()
}

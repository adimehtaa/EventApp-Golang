package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		log.Fatal("Plese provie a migration direction 'up' or 'down' ")
	}
}

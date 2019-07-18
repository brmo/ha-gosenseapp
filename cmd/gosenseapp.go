package main

import (
	"fmt"
	"os"

	"github.com/brmo/ha-gosenseapp"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("ha-gosenseapp starting...")

	log.SetLevel(log.InfoLevel)
	log.SetOutput(os.Stdout)

	gosenseapp.Run()
}

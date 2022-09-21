package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/synthonier/me-sniper/pkg/sniper"
)

func main() {
	err := godotenv.Load()
	checkError(err)

	s, err := sniper.New(os.Getenv("NODE_ENDPOINT"))
	checkError(err)
	
	err = s.Start()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// R, _ := Initrdrr(os.Getenv("RDRRKEY"), "http://localhost:7878")
	// R.getQualityProfile()
	// R.getRootFolder()
	// N := notion.InitNotionClient("Emad", os.Getenv("RDRRNOTIONINTEG"), os.Getenv("DBID"))
	// data, err := N.QueryDB()

}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
}

func GetTwitterApi() *anaconda.TwitterApi {
	return anaconda.NewTwitterApiWithCredentials(
		os.Getenv("accessToken"),
		os.Getenv("accessSecret"),
		os.Getenv("consumerKey"),
		os.Getenv("consumerSecret"),
	)
}

func main() {
	LoadEnv()
	api := GetTwitterApi()
	flag.Parse()

	_, err := api.PostTweet(flag.Arg(0), nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Tweet completed.")
}

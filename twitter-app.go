package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"os"
)

func main() {
	anaconda.SetConsumerKey("")
	anaconda.SetConsumerSecret("")

	api := anaconda.NewTwitterApi("", "")

	tw := "Test post please ignore"

	tweet, err := api.PostTweet(tw, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(tweet.Text)

}

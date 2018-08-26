package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/daveilers/everylotbot/everylot"
)

var iD = flag.String(`id`, "", `tweet the entry in the lots table with this id`)

func main() {
	flag.Parse()
	var l *everylot.Lot
	var err error

	if iD == nil || *iD == "" {
		l, err = everylot.Next()
	} else {
		l, err = everylot.ID(*iD)
	}
	if err != nil {
		log.Fatalf("Error getting lot: %v", err)
	}

	tweetId, err := l.PostTweet()
	if err != nil {
		log.Fatalf("Broke attempting to tweet %v: %v", l, err)
	}
	l.Tweeted = tweetId

	err = l.MarkAsTweeted(tweetId)
	if err != nil {
		log.Fatalf("Broke attempting to mark as tweeted tweet %v: %v", l, err)
	}

	fmt.Printf("Complete success tweeting %#v\n", l)

}

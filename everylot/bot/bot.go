package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/daveilers/everylotbot/everylot"
	"github.com/daveilers/everylotbot/everylot/db"
)

var iD = flag.String(`id`, "", `tweet the entry in the lots table with this id`)

func main() {
	flag.Parse()
	var l *everylot.Lot
	var err error

	db, err := db.NewSQLite(os.Getenv("DB_DRIVER"), os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatalf("Error getting lot: %v", err)
	}

	if iD == nil || *iD == "" {
		l, err = db.Next()
	} else {
		l, err = db.ID(*iD)
	}
	if err != nil {
		log.Fatalf("Error getting lot: %v", err)
	}

	tweetId, err := l.PostTweet()
	if err != nil {
		err = db.MarkAsTweeted(l, "-1")
		log.Fatalf("Broke attempting to tweet %v: %v", l, err)
	}
	l.Tweeted = tweetId

	err = db.MarkAsTweeted(l, tweetId)
	if err != nil {
		log.Fatalf("Broke attempting to mark as tweeted tweet %v: %v", l, err)
	}

	fmt.Printf("Complete success tweeting %#v\n", l)

}

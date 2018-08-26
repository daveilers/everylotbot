package everylot

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var svc *Service

func init() {
	svc = &Service{
		apiKey:            os.Getenv("GMAPS"),
		AppName:           os.Getenv("APPNAME"),
		AppConsumerKey:    os.Getenv("APP_CONSUMER_KEY"),
		AppConsumerSecret: os.Getenv("APP_CONSUMER_SECRET"),
		UserName:          os.Getenv("USER"),
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
	}
}

func Next() {
	// return next non-empty parcel number's streetview tweet

	//SELECT * FROM lots WHERE tweeted = 0 ORDER BY id LIMIT 1;

}

func ID(parcelno int) {
	// id, err := lot.PostTweet()
	// if err != nil {
	// 	log.Fatalf("Broke: %v", err)
	// }
	// fmt.Println(id)
}

var debug = &log.Logger{}
var info = &log.Logger{}

func (l *Lot) AddressString() string {
	return fmt.Sprintf("%s, %s %s", l.Address, l.City, l.State)
}

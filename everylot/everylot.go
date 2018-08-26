package everylot

import (
	"fmt"
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
		DBDriver:          os.Getenv("DB_DRIVER"),
		DBDSN:             os.Getenv("DB_DSN"),
	}
}

// var debug = &log.Logger{}
// var info = &log.Logger{}

func (l *Lot) AddressString() string {
	if l == nil {
		panic("That's not a lot")
	}
	return fmt.Sprintf("%s, %s %s", l.Address, l.City, l.State)
}

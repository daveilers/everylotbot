package everylot

type Service struct {
	apiKey            string
	AppName           string
	AppConsumerKey    string
	AppConsumerSecret string
	UserName          string
	AccessToken       string
	AccessTokenSecret string
	DBDriver          string
	DBDSN             string
}

type Lot struct {
	ID           string
	Lat          float64
	Lon          float64
	Address      string
	City         string
	State        string
	Zip          string
	GMapsID      string
	GMapsAddress string
	Tweeted      string
}

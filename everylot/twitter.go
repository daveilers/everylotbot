package everylot

import (
	"encoding/base64"
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
)

func (svc *Service) twitInit() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(svc.AppConsumerKey)
	anaconda.SetConsumerSecret(svc.AppConsumerSecret)
	api := anaconda.NewTwitterApi(svc.AccessToken, svc.AccessTokenSecret)
	return api
}

func (l *Lot) PostTweet() (tweetId string, err error) {
	api := svc.twitInit()
	d, err := l.getStreetviewImage()
	if err != nil {
		return
	}
	m, err := api.UploadMedia(base64.StdEncoding.EncodeToString(d))
	if err != nil {
		return
	}
	tw, err := api.PostTweet(l.Address, url.Values{
		"lat":       []string{strconv.FormatFloat(l.Lat, 'f', -1, 64)},
		"long":      []string{strconv.FormatFloat(l.Lon, 'f', -1, 64)},
		"media_ids": []string{m.MediaIDString},
	})
	return tw.IdStr, err
}

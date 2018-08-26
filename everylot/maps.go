package everylot

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"googlemaps.github.io/maps"
)

const sVAPI = "https://maps.googleapis.com/maps/api/streetview"
const gCAPI = "https://maps.googleapis.com/maps/api/geocode/json"

func (l *Lot) getStreetviewImage() (data []byte, err error) {
	// Fetch image from streetview API

	params := url.Values{
		"location": []string{l.streetviewableLocation()},
		"size":     []string{"1000x1000"},
		"fov":      []string{"65"},
		"pitch":    []string{"10"},
		"key":      []string{svc.apiKey},
	}

	u, err := url.Parse(sVAPI)
	if err != nil {
		return
	}
	u.RawQuery = params.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)

}

func (l *Lot) streetviewableLocation() string {
	// search google for lot address
	err := l.geoCodeFromAddress()
	if err != nil {
		debug.Printf("%v", err)
	}

	// if we had lat&long we could use them here to doublecheck or instead, but we don't so oh well.
	return l.AddressString()
}

func (l *Lot) geoCodeFromAddress() (err error) {
	c, err := maps.NewClient(maps.WithAPIKey(svc.apiKey))
	if err != nil {
		return err
	}
	results, err := c.Geocode(context.Background(), &maps.GeocodingRequest{Address: l.AddressString()})
	if err != nil {
		return
	}
	if len(results) == 0 {
		return fmt.Errorf("No result found")
	}
	result := results[0]
	if result.PartialMatch {
		return fmt.Errorf("Imperfect match")
	}
	l.GMapsAddress = result.FormattedAddress
	l.Lat = result.Geometry.Location.Lat
	l.Lon = result.Geometry.Location.Lng
	l.GMapsID = result.PlaceID
	return nil
}

//     def aim_camera(self):
//         '''Set field-of-view and pitch'''
//         fov, pitch = 65, 10
//         try:
//             floors = float(self.lot.get('floors', 0)) or 2
//         except TypeError:
//             floors = 2

//         if floors == 3:
//             fov = 72

//         if floors == 4:
//             fov, pitch = 76, 15

//         if floors >= 5:
//             fov, pitch = 81, 20

//         if floors == 6:
//             fov = 86

//         if floors >= 8:
//             fov, pitch = 90, 25

//         if floors >= 10:
//             fov, pitch = 90, 30

//         return fov, pitch

package latlong

import "fmt"

type Place struct {
	latitude, longitude float64
	Name                string
}

func New(latitude, longitude float64, name string) *Place {
	return &Place{
		latitude, longitude, name,
	}
}

func (place *Place) Latitude() float64 {
	return place.latitude
}

func (place *Place) Longitude() float64 {
	return place.longitude
}

func (place *Place) SetLatitude(latitude float64) {
	place.latitude = latitude
}

func (place *Place) SetLongitude(longitude float64) {
	place.longitude = longitude
}

func (place *Place) String() string {
	return fmt.Sprintf("(%.3f°, %.3f°) %q", place.latitude,
		place.longitude, place.Name)
}

func (place *Place) Copy() *Place {
	return &Place{
		place.latitude,
		place.longitude,
		place.Name,
	}
}
package utils

import (
	"math"
	"time"

	"golang.org/x/exp/rand"
)

const EARTH_RADIUS = 6371009 // meters
const DEG_TO_RAD = math.Pi / 180.0
const THREE_PI = math.Pi * 3
const TWO_PI = math.Pi * 2

func toRadians(deg float64) float64 {
	return deg * DEG_TO_RAD
}

func toDegrees(rad float64) float64 {
	return rad / DEG_TO_RAD
}

var rng = rand.New(rand.NewSource(uint64(time.Now().UTC().UnixNano())))

// random returns a float64 between 0 and 1
func random() float64 {
	return rng.Float64()
}

type Point struct {
	Latitude  float64
	Longitude float64
}

/*
Given a centerPoint C and a radius R, returns a random point that is on the
circumference defined by C and R.

centerPoint C is of type { latitude: A, longitude: B }
Where -90 <= A <= 90 and -180 <= B <= 180.

radius R is in meters.

Based on: http://www.movable-type.co.uk/scripts/latlong.html#destPoint
*/
func RandomCircumferencePoint(centerPoint Point, radius float64) Point {

	sinLat := math.Sin(toRadians(centerPoint.Latitude))
	cosLat := math.Cos(toRadians(centerPoint.Latitude))

	// Random bearing (direction out 360 degrees)
	bearing := random() * TWO_PI
	sinBearing := math.Sin(bearing)
	cosBearing := math.Cos(bearing)

	// Theta is the approximated angular distance
	theta := radius / EARTH_RADIUS
	sinTheta := math.Sin(theta)
	cosTheta := math.Cos(theta)

	rLatitude := math.Asin(sinLat*cosTheta + cosLat*sinTheta*cosBearing)
	rLongitude := toRadians(centerPoint.Longitude) +
		math.Atan2(sinBearing*sinTheta*cosLat, cosTheta-sinLat*math.Sin(rLatitude))

	// Normalize longitude L such that -PI < L < +PI
	rLongitude = math.Mod(rLongitude+THREE_PI, TWO_PI) - math.Pi

	return Point{
		Latitude:  toDegrees(rLatitude),
		Longitude: toDegrees(rLongitude),
	}
}

/*
Given a centerPoint C and a radius R, returns a random point that is inside
the circle defined by C and R.

centerPoint C is of type { latitude: A, longitude: B }
Where -90 <= A <= 90 and -180 <= B <= 180.

radius R is in meters.
*/
func RandomCirclePoint(centerPoint Point, radius float64) Point {
	return RandomCircumferencePoint(centerPoint, math.Sqrt(random())*radius)
}

// radius in meters
func GenerateRandomLatLon(lat, lon float64, radiusMin int, radiusMax int) Point {
	radius := rand.Intn(radiusMax-radiusMin) + radiusMin
	return RandomCirclePoint(Point{Latitude: lat, Longitude: lon}, float64(radius))
}

package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func SecondHand(t time.Time) Point {
	p := SecondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}
	return p
}

func SecondHandPoint(tm time.Time) Point {
	return angleToPoint(SecondsInRadians(tm))
}

func SecondsInRadians(tm time.Time) float64 {
	return (math.Pi / (30 / (float64(tm.Second()))))
}

func minutesInRadians(t time.Time) float64 {
	return (SecondsInRadians(t) / 60) +
		(math.Pi / (30 / float64(t.Minute())))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / 12) +
		(math.Pi / (6 / float64(t.Hour()%12)))
}

func MinuteHandPoint(tm time.Time) Point {
	return angleToPoint(minutesInRadians(tm))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

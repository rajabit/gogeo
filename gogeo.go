package gogeo

import (
	"math"

	"github.com/chewxy/math32"
)

type (
	Point struct {
		Latitude  float32
		Longitude float32
	}

	Polygon struct {
		Points []*Point
	}
)

func NewPolygon(in [][]float32) *Polygon {
	p := &Polygon{}
	for _, x := range in {
		p.NewPoint(NewPoint(x[1], x[0]))
	}
	return p
}

func NewPoint(lat float32, lng float32) *Point {
	return &Point{
		Latitude:  lat,
		Longitude: lng,
	}
}

func (p *Polygon) RectangleContains(point *Point) bool {
	minLat, maxLat, minLng, maxLng := p.BindingPoint()
	return point.Latitude >= minLat && point.Latitude <= maxLat && point.Longitude >= minLng && point.Longitude <= maxLng
}

func (p *Polygon) PolygonContains(point *Point) bool {
	if !p.IsClosed() {
		return false
	}

	start := len(p.Points) - 1
	end := 0

	contains := p.rayCast(point, p.Points[start], p.Points[end])

	for i := 1; i < len(p.Points); i++ {
		if p.rayCast(point, p.Points[i-1], p.Points[i]) {
			contains = !contains
		}
	}

	return contains
}

func (p *Polygon) CircleContains(point *Point, radius float32) bool {
	distance := p.Points[0].DistanceInMeter(point)
	return distance <= radius
}

func (p *Polygon) BindingPoint() (minLat, maxLat, minLng, maxLng float32) {
	minLat, maxLat = math.MaxFloat32, -math.MaxFloat32
	minLng, maxLng = math.MaxFloat32, -math.MaxFloat32

	for _, point := range p.Points {
		if point.Latitude < minLat {
			minLat = point.Latitude
		}
		if point.Latitude > maxLat {
			maxLat = point.Latitude
		}
		if point.Longitude < minLng {
			minLng = point.Longitude
		}
		if point.Longitude > maxLng {
			maxLng = point.Longitude
		}
	}

	return
}

func (p *Polygon) rayCast(point *Point, start *Point, end *Point) bool {
	if start.Longitude > end.Longitude {
		start, end = end, start

	}

	for point.Longitude == start.Longitude || point.Longitude == end.Longitude {
		newLng := math32.Nextafter(point.Longitude, math32.Inf(1))
		point = NewPoint(point.Latitude, newLng)
	}

	if point.Longitude < start.Longitude || point.Longitude > end.Longitude {
		return false
	}

	if start.Latitude > end.Latitude {
		if point.Latitude > start.Latitude {
			return false
		}
		if point.Latitude < end.Latitude {
			return true
		}

	} else {
		if point.Latitude > end.Latitude {
			return false
		}
		if point.Latitude < start.Latitude {
			return true
		}
	}

	raySlope := (point.Longitude - start.Longitude) / (point.Latitude - start.Latitude)
	diagSlope := (end.Longitude - start.Longitude) / (end.Latitude - start.Latitude)

	return raySlope >= diagSlope
}

func (p *Polygon) IsClosed() bool {
	return !(len(p.Points) < 3)
}

func (p *Polygon) NewPoint(point *Point) {
	p.Points = append(p.Points, point)
}

func (p *Point) DistanceInMeter(point *Point) float32 {
	var la1, lo1, la2, lo2 float32
	piRad := float32(math.Pi) / 180
	la1 = p.Latitude * piRad
	lo1 = p.Longitude * piRad
	la2 = point.Latitude * piRad
	lo2 = point.Longitude * piRad
	h := hsin(la2-la1) + math32.Cos(la1)*math32.Cos(la2)*hsin(lo2-lo1)
	meters := 2 * 6378100 * math32.Asin(math32.Sqrt(h))
	return meters
}

func hsin(theta float32) float32 {
	return math32.Pow(math32.Sin(theta)/2, 2)
}

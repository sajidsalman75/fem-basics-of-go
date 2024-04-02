package data

type distance float64 // miles
type distanceKm float64
type integer = float64

// Method
func (miles distance) ToKm() distanceKm {
	return distanceKm(miles * 1.60934)
}

func (km distanceKm) ToMiles() distance {
	return distance(km / 1.60934)
}

func test() {

	d := distance(4.5)
	km := d.ToKm()
	miles := km.ToMiles()

	print(miles)
}

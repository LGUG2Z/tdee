package cmd

const (
	FeetToCentimetres float64 = 30.48
	PoundsToKilograms float64 = 0.453592
)

type Units int

const (
	Imperial Units = iota
	Metric
)

const (
	Female = iota
	Male
)

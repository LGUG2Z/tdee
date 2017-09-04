package cmd

import (
	"fmt"
	"strings"
)

func determineUnits(imperial, metric bool) (Units, error) {
	if !imperial && !metric {
		return Units(0), ErrNoUnitsSelected
	}

	if imperial {
		return Imperial, nil
	}

	return Metric, nil
}

func hasRequiredInformation(height, weight, age, lifestyle float64, sex string) bool {
	if height > 0 && weight > 0 && age > 0 && lifestyle > 0 && len(sex) > 0 {
		return true
	}

	return false
}

func hasValidLifestyleModifier(l float64) bool {
	valid := []float64{1.2, 1.375, 1.55, 1.7, 1.9}

	for _, v := range valid {
		if l == v {
			return true
		}
	}

	return false
}

type Information struct {
	Height    float64
	Weight    float64
	Age       float64
	Sex       int
	Lifestyle float64
}

func (i *Information) FromInput(height, weight, age, lifestyle float64, sex string) error {
	i.Height = height
	i.Weight = weight
	i.Age = age

	if strings.ToLower(sex) == "male" {
		i.Sex = Male
	} else if strings.ToLower(sex) == "female" {
		i.Sex = Female
	} else {
		return ErrInvalidSex
	}

	if !hasValidLifestyleModifier(lifestyle) {
		return ErrInvalidLifestyleModifier
	}

	i.Lifestyle = lifestyle

	return nil
}

func (i *Information) CalculateTDEE(u Units) string {
	msj := i.MifflinStJeor(u)
	hbo := i.HarrisBenedictOriginal(u)
	hbr := i.HarrisBenedictRevised(u)

	avg := (msj + hbo + hbr) / 3
	tdee := avg * i.Lifestyle

	return fmt.Sprintf("%.0f", tdee)
}

func (i *Information) MifflinStJeor(u Units) float64 {
	if u == Imperial {
		var m Information
		m.Lifestyle = i.Lifestyle
		m.Sex = i.Sex
		m.Age = i.Age

		m.Height = i.Height * FeetToCentimetres
		m.Weight = i.Weight * PoundsToKilograms

		return m.MifflinStJeor(Metric)
	}

	weightMultiplier := 10.0
	heightMultiplier := 6.25
	ageMultiplier := 5.0
	maleAdjustment := 5.0
	femaleAdjustment := 161.0

	unadjustedCalculation := (i.Weight * weightMultiplier) + (i.Height * heightMultiplier) - (i.Age * ageMultiplier)

	if i.Sex == Female {
		return unadjustedCalculation - femaleAdjustment
	}

	// Else male
	return unadjustedCalculation + maleAdjustment
}

func (i *Information) HarrisBenedictOriginal(u Units) float64 {
	if u == Imperial {
		var m Information
		m.Lifestyle = i.Lifestyle
		m.Sex = i.Sex
		m.Age = i.Age

		m.Height = i.Height * FeetToCentimetres
		m.Weight = i.Weight * PoundsToKilograms

		return m.HarrisBenedictOriginal(Metric)
	}

	var weightMultiplier float64
	var heightMultiplier float64
	var ageMultiplier float64
	var unadjustedCalculation float64

	if i.Sex == Female {
		weightMultiplier = 9.5643
		heightMultiplier = 1.8496
		ageMultiplier = 4.6756
		femaleAdjustment := 655.0955

		unadjustedCalculation = (i.Weight * weightMultiplier) + (i.Height * heightMultiplier) - (i.Age * ageMultiplier)
		return unadjustedCalculation + femaleAdjustment
	}

	// Else male
	weightMultiplier = 13.7516
	heightMultiplier = 5.0033
	ageMultiplier = 6.7550
	maleAdjustment := 66.4730

	unadjustedCalculation = (i.Weight * weightMultiplier) + (i.Height * heightMultiplier) - (i.Age * ageMultiplier)
	return unadjustedCalculation + maleAdjustment

}

func (i *Information) HarrisBenedictRevised(u Units) float64 {
	if u == Imperial {
		var m Information
		m.Lifestyle = i.Lifestyle
		m.Sex = i.Sex
		m.Age = i.Age

		m.Height = i.Height * FeetToCentimetres
		m.Weight = i.Weight * PoundsToKilograms

		return m.HarrisBenedictRevised(Metric)
	}

	var weightMultiplier float64
	var heightMultiplier float64
	var ageMultiplier float64
	var unadjustedCalculation float64

	if i.Sex == Female {
		weightMultiplier = 9.247
		heightMultiplier = 3.098
		ageMultiplier = 4.330
		femaleAdjustment := 447.593

		unadjustedCalculation = (i.Weight * weightMultiplier) + (i.Height * heightMultiplier) - (i.Age * ageMultiplier)
		return unadjustedCalculation + femaleAdjustment
	}

	// Else male
	weightMultiplier = 13.379
	heightMultiplier = 4.799
	ageMultiplier = 5.677
	maleAdjustment := 88.362

	unadjustedCalculation = (i.Weight * weightMultiplier) + (i.Height * heightMultiplier) - (i.Age * ageMultiplier)
	return unadjustedCalculation + maleAdjustment
}

package cmd

import "errors"

var (
	ErrMissingInformation       = errors.New("Not all information has been provided.")
	ErrNoUnitsSelected          = errors.New("Imperial or metric units must be selected.")
	ErrInvalidLifestyleModifier = errors.New("Invalid lifestyle modifier provided.")
	ErrInvalidSex               = errors.New("Sex must be provided as male or female.")
)

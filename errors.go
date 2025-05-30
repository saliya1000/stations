package main

import "fmt"

type ValidationError struct {
	Code    string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func NewValidationError(code string, args ...string) error {
	return &ValidationError{
		Code:    code,
		Message: fmt.Sprintf("%s: %v", code, args),
	}
}

func ValidateStationExists(network *Network, station string, isStart bool) error {
	if network.GetStationByName(station) == nil {
		if isStart {
			return NewValidationError("invalid start station", station)
		}
		return NewValidationError("invalid end station", station)
	}
	return nil
}


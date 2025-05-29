package main

import (
	"os"
	"strings"
	"testing"
)

func TestParseNetworkFile_DuplicateStationName(t *testing.T) {
	content := `
stations:
alpha,1,1
beta,2,2
alpha,3,3
connections:
alpha-beta
`
	tmpfile, err := os.CreateTemp("", "test_network_*.map")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	_, err = ParseNetworkFile(tmpfile.Name())
	if err == nil {
		t.Errorf("Expected an error for duplicate station name, but got nil")
		return
	}

	validationErr, ok := err.(*ValidationError)
	if !ok {
		t.Errorf("Expected error to be of type *ValidationError, but got %T: %v", err, err)
		return
	}

	expectedErrorCode := "duplicate_station_name"
	if validationErr.Code != expectedErrorCode {
		t.Errorf("Expected error code %s, but got %s", expectedErrorCode, validationErr.Code)
	}

	// Test for the specific station name in the error message
	// The NewValidationError formats it as "code: [args]"
	expectedStationName := "alpha"
	// As per errors.go NewValidationError, Message is: fmt.Sprintf("%s: %v", code, args)
	// So for a single arg, it would be "duplicate_station_name: [alpha]"
	expectedMessageDetail := "[" + expectedStationName + "]" 

	if !strings.Contains(validationErr.Message, expectedErrorCode) {
		t.Errorf("Expected error message to contain code '%s', but got '%s'", expectedErrorCode, validationErr.Message)
	}
	if !strings.Contains(validationErr.Message, expectedMessageDetail) {
		t.Errorf("Expected error message to contain station detail '%s', but got '%s'", expectedMessageDetail, validationErr.Message)
	}
}

func TestParseNetworkFile_DuplicateConnection(t *testing.T) {
	cases := []struct {
		name             string
		content          string
		expectedStations []string // Stations involved in the duplicate connection
	}{
		{
			name: "explicit duplicate connection",
			content: `
stations:
alpha,1,1
beta,2,2
gamma,3,3
connections:
alpha-beta
gamma-alpha
alpha-beta
`,
			expectedStations: []string{"alpha", "beta"},
		},
		{
			name: "implicit duplicate connection",
			content: `
stations:
alpha,1,1
beta,2,2
gamma,3,3
connections:
alpha-beta
gamma-alpha
beta-alpha
`,
			expectedStations: []string{"alpha", "beta"},
		},
		{
			name: "implicit duplicate connection different order",
			content: `
stations:
delta,1,1
epsilon,2,2
zeta,3,3
connections:
epsilon-delta
zeta-epsilon
delta-epsilon
`,
			expectedStations: []string{"delta", "epsilon"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tmpfile, err := os.CreateTemp("", "test_network_*.map")
			if err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}
			defer os.Remove(tmpfile.Name()) // clean up

			if _, err := tmpfile.WriteString(tc.content); err != nil {
				t.Fatalf("Failed to write to temp file: %v", err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatalf("Failed to close temp file: %v", err)
			}

			_, err = ParseNetworkFile(tmpfile.Name())
			if err == nil {
				t.Errorf("Expected an error for duplicate connection, but got nil")
				return
			}

			validationErr, ok := err.(*ValidationError)
			if !ok {
				t.Errorf("Expected error to be of type *ValidationError, but got %T: %v", err, err)
				return
			}

			expectedErrorCode := "duplicate_connection"
			if validationErr.Code != expectedErrorCode {
				t.Errorf("Expected error code %s, but got %s", expectedErrorCode, validationErr.Code)
			}

			// Check if the error message contains the error code
			if !strings.Contains(validationErr.Message, expectedErrorCode) {
				t.Errorf("Expected error message to contain code '%s', but got '%s'", expectedErrorCode, validationErr.Message)
			}

			// Check if the error message contains both station names
			// Note: The order of stations in the message might vary depending on how they were processed (e.g. "alpha beta" or "beta alpha")
			// So we check for the presence of each station name individually.
			// The NewValidationError formats it as "code: [arg1 arg2 ...]"
			// Example: "duplicate_connection: [alpha beta]"
			for _, stationName := range tc.expectedStations {
				if !strings.Contains(validationErr.Message, stationName) {
					t.Errorf("Expected error message to contain station name '%s', but got '%s'", stationName, validationErr.Message)
				}
			}
		})
	}
}

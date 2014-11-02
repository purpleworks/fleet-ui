package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// GetMachineIP parses the unitMachine in format "uuid/ip" and returns only the IP part.
// Can be used with the {UnitStatus.Machine} field.
// Returns an empty string, if no ip was found.
func GetMachineIP(unitMachine string) string {
	fields := strings.Split(unitMachine, "/")
	if len(fields) < 2 {
		return ""
	}
	return fields[1]
}

func execCmd(cmd *exec.Cmd) (string, error) {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if result := stdout.String(); result != "" {
		return result, nil
	}

	if err != nil {
		return "", err
	}

	if err := stderr.String(); err != "" {
		return "", fmt.Errorf(err)
	}

	return "", nil
}

// filterEmpty returns an array containing all non-empty strings of the input array.
// Non-empty as in `strings.TrimSpace(v) != ""`.
func filterEmpty(values []string) []string {
	result := make([]string, 0)
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			result = append(result, v)
		}
	}
	return result
}

// error handling

const (
	ERROR_TYPE_NOT_FOUND = 10000 + iota
)

type FleetClientError struct {
	StatusCode int
	StatusText string
}

func (this FleetClientError) Error() string {
	return this.StatusText
}

func NewFleetClientError(code int, text string) FleetClientError {
	return FleetClientError{
		StatusCode: code,
		StatusText: text,
	}
}

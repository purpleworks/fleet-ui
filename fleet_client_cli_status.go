package main

// Parse for `fleet status` output
import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func (this *ClientCLI) StatusAll() ([]UnitStatus, error) {
	cmd := exec.Command(FLEETCTL, ENDPOINT_OPTION, this.etcdPeer, "list-units", "--full=true", "-l=true", "--fields=unit,load,active,sub,machine")
	stdout, err := execCmd(cmd)
	if err != nil {
		return []UnitStatus{}, err
	}

	return parseFleetStatusOutput(stdout)
}

func parseFleetStatusOutput(output string) ([]UnitStatus, error) {
	result := make([]UnitStatus, 0)

	scanner := bufio.NewScanner(strings.NewReader(output))
	// Scan each line of input.
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		if lineCount == 1 {
			continue
		}

		words := filterEmpty(strings.Split(line, "\t"))
		unitStatus := UnitStatus{
			Unit:    words[0],
			Load:    words[1],
			Active:  words[2],
			Sub:     words[3],
			Machine: words[4],
		}
		result = append(result, unitStatus)
	}

	if err := scanner.Err(); err != nil {
		return result, scanner.Err()
	}
	return result, nil
}

func (this *ClientCLI) StatusUnit(name string) (UnitStatus, error) {
	status, err := this.StatusAll()
	if err != nil {
		return UnitStatus{}, err
	}

	for _, s := range status {
		if s.Unit == name {
			return s, nil
		}
	}

	return UnitStatus{}, NewFleetClientError(ERROR_TYPE_NOT_FOUND, fmt.Sprintf("Cannot fetch status. Unit (%s) not found. Aborting...", name))
}

func (this *ClientCLI) JournalF(name string) (outc, errc chan string, err error) {
	var stdout, stderr io.ReadCloser

	cmdY := exec.Command("echo", "y")
	cmd := exec.Command(FLEETCTL, ENDPOINT_OPTION, this.etcdPeer, "journal", "-f", name)

	cmd.Stdin, _ = cmdY.StdoutPipe()
	stdout, err = cmd.StdoutPipe()
	if err != nil {
		return
	}

	stderr, err = cmd.StderrPipe()
	if err != nil {
		return
	}

	if err = cmd.Start(); err != nil {
		return
	}

	cmdY.Run()

	outc = make(chan string)
	errc = make(chan string)

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			outc <- scanner.Text()
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			errc <- scanner.Text()
		}
	}()

	return
}

func (this *ClientCLI) MachineAll() ([]MachineStatus, error) {
	cmd := exec.Command(FLEETCTL, ENDPOINT_OPTION, this.etcdPeer, "list-machines", "--full=true")
	stdout, err := execCmd(cmd)
	if err != nil {
		return []MachineStatus{}, err
	}

	return parseMachineStatusOutput(stdout)
}

func parseMachineStatusOutput(output string) ([]MachineStatus, error) {
	result := make([]MachineStatus, 0)

	scanner := bufio.NewScanner(strings.NewReader(output))
	// Scan each line of input.
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		if lineCount == 1 {
			continue
		}

		words := filterEmpty(strings.Split(line, "\t"))
		unitStatus := MachineStatus{
			Machine:   words[0],
			IPAddress: words[1],
			Metadata:  words[2],
		}
		result = append(result, unitStatus)
	}

	if err := scanner.Err(); err != nil {
		return result, scanner.Err()
	}
	return result, nil
}

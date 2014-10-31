package main

// Parse for `fleet status` output
import (
  "bufio"
  "fmt"
  "strings"

  execPkg "os/exec"
)

// MachineIP returns the IP of the used machine, if available, an empty string otherwise.
// Alias for `return GetMachineIp(status.Machine)`.
func (status UnitStatus) MachineIP() string {
  return GetMachineIP(status.Machine)
}

// StatusAll executes "fleetctl status" and parses the output table. Thus, certain fields can be mangled or
// shortened, e.g. the machine column.
func (this *ClientCLI) StatusAll() ([]UnitStatus, error) {
  cmd := execPkg.Command(FLEETCTL, ENDPOINT_OPTION, this.etcdPeer, "list-units", "--full=true", "-l=true", "--fields=unit,load,active,sub,machine")
  stdout, err := exec(cmd)
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

  // When finished scanning if any error other than io.EOF occured
  // it will be returned by scanner.Err().
  if err := scanner.Err(); err != nil {
    return result, scanner.Err()
  }
  return result, nil
}

// StatusUnit returns the UnitStatus for the given unitfile.
// If the unit is not found or could not be retrieved, an error will be returned.
//
// Internally it executes `fleetctl status` and looks for the matching
// unit file. Its {UnitStatus} will be returned.
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

func (this *ClientCLI) Status(name string) (*Status, error) {
  allStatus, err := this.StatusAll()
  if err != nil {
    return nil, err
  }

  for _, status := range allStatus {
    if status.Unit == name {
      return &Status{
        Running:     status.Sub == SUB_RUNNING,
        ContainerIP: "127.0.0.1",
      }, nil
    }
  }
  // Return running=false, because we didn't find it
  return nil, fmt.Errorf("Job not found: %s", name)
}

func (this *ClientCLI) JournalF(name string) (chan string, error) {
  cmdY := execPkg.Command("echo", "y")
  cmd := execPkg.Command(FLEETCTL, ENDPOINT_OPTION, this.etcdPeer, "journal", "-f", name)

  cmd.Stdin, _ = cmdY.StdoutPipe()
  stdout, err := cmd.StdoutPipe()
  if err != nil {
    return nil, err
  }
  if err := cmd.Start(); err != nil {
    return nil, err
  }
  cmdY.Run()

  linec := make(chan string)
  scanner := bufio.NewScanner(stdout)
  go func() {
    for scanner.Scan() {
      linec <- scanner.Text()
    }
  }()
  return linec, nil
}

func (this *ClientCLI) MachineAll() ([]MachineStatus, error) {
  cmd := execPkg.Command(FLEETCTL, ENDPOINT_OPTION, this.etcdPeer, "list-machines", "--full=true")
  stdout, err := exec(cmd)
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

  // When finished scanning if any error other than io.EOF occured
  // it will be returned by scanner.Err().
  if err := scanner.Err(); err != nil {
    return result, scanner.Err()
  }
  return result, nil
}

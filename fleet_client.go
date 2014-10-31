package main

import (
  "github.com/coreos/fleet/schema"
)

const (
  STATE_LAUNCHED = "launched"
  STATE_LOADED   = "loaded"
  STATE_INACTIVE = "inactive"

  LOAD_UNKNOWN = "-"
  LOAD_LOADED  = "loaded" // See https://github.com/coreos/fleet/blob/master/job/job.go

  ACTIVE_UNKNOWN    = "-"
  ACTIVE_ACTIVE     = "active"
  ACTIVE_ACTIVATING = "activating"
  ACTIVE_FAILED     = "failed"

  SUB_UNKNOWN   = "-"
  SUB_DEAD      = "dead"
  SUB_LAUNCHED  = "launched"
  SUB_START     = "start"
  SUB_START_PRE = "start-pre"
  SUB_RUNNING   = "running"
  SUB_EXITED    = "exited"
  SUB_FAILED    = "failed"
)

type UnitStatus struct {
  // Unit Name with file extension
  Unit string

  // Fleet state, "launched" or "inactive"
  State string

  // "-", "loaded"
  Load string

  // "-", "active", "failed"
  Active string

  // The state of the unit, e.g. "-", "running" or "failed". See the SUB_* constants.
  Sub string

  Description string

  // The machine that is used to execute the unit.
  // Is "-", when no machine is assigned.
  // Otherwise is in the format of "uuid/ip", where uuid is shortened version of the host uuid
  // and IP is the IP assigned to that machine.
  Machine string
}

type MachineStatus struct {
  Machine   string
  IPAddress string
  Metadata  string
}

type Status struct {
  Running     bool
  ContainerIP string
}

type FleetClient interface {
  // A Unit is a submitted job known by fleet, but not started yet. Submitting
  // a job creates a unit. Unit() returns such an object. Further a Unit has
  // different properties than a ScheduledUnit.
  Unit(name string) (*schema.Unit, error)

  Submit(name, filePath string) error
  Start(name string) error
  Stop(name string) error
  Load(name string) error
  Destroy(name string) error
  Status(name string) (*Status, error) // Deprecated, use StatusUnit()
  StatusUnit(name string) (UnitStatus, error)
  StatusAll() ([]UnitStatus, error)
  JournalF(name string) (chan string, error)
  MachineAll() ([]MachineStatus, error)
}

func NewClient() FleetClient {
  return NewClientCLI()
}

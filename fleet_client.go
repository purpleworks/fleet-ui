package main

type UnitStatus struct {
	Unit, Load, Active, Sub, Machine string
}

type MachineStatus struct {
	Machine, IPAddress, Metadata string
}

type Status struct {
	Running     bool
	ContainerIP string
}

type FleetClient interface {
	Submit(name, filePath string) error
	Start(name string) error
	Stop(name string) error
	Load(name string) error
	Destroy(name string) error
	StatusUnit(name string) (UnitStatus, error)
	StatusAll() ([]UnitStatus, error)
	JournalF(name string) (chan string, chan string, error)
	MachineAll() ([]MachineStatus, error)
}

func NewClient() FleetClient {
	return NewClientCLI()
}

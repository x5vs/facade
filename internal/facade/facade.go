package facade

import (
	"github.com/x5vs/facade/internal/facade/cpu"
	"github.com/x5vs/facade/internal/facade/drive"
	"github.com/x5vs/facade/internal/facade/memory"
)

// Computer интерфейс камьютера
type Computer interface {
	Start()
}

type computer struct {
	bootAddress       uint64
	bootSectorAddress uint64
	bootSectorSize    uint64
	cpu               cpu.Proccessor
	memory            memory.Memory
	drive             drive.Drive
}

// Start a computer
func (c *computer) Start() {
	c.cpu.Freeze()

	data, err := c.drive.Read(c.bootSectorAddress, c.bootSectorSize)
	if err != nil {
		panic(err)
	}

	err = c.memory.Load(c.bootAddress, data)
	if err != nil {
		panic(err)
	}

	c.cpu.Jump(c.bootAddress)
	c.cpu.Execute()
}

// NewComputer create new computer
func NewComputer(
	cpu cpu.Proccessor,
	drive drive.Drive,
	memory memory.Memory,
) Computer {
	return &computer{
		bootAddress:       0,
		bootSectorAddress: 0,
		bootSectorSize:    128,
		cpu:               cpu,
		drive:             drive,
		memory:            memory,
	}
}

package facade

import (
	"context"
)

// Computer интерфейс камьютера
type Computer interface {
	// Start a computer
	Start(context.Context) error
}

type cpu interface {
	// Freeze
	Freeze(context.Context) error
	// Jump to address
	Jump(context.Context, uint64) error
	// Execute command
	Execute(context.Context) error
}

type memory interface {
	// Load data to memory
	Load(ctx context.Context, position uint64, data []byte) error
}

type drive interface {
	// Read data from address
	Read(ctx context.Context, lba, size uint64) ([]byte, error)
}

type computer struct {
	bootAddress       uint64
	bootSectorAddress uint64
	bootSectorSize    uint64
	cpu               cpu
	memory            memory
	drive             drive
}

// Start a computer
func (c *computer) Start(ctx context.Context) error {
	err := c.cpu.Freeze(ctx)
	if err != nil {
		return err
	}
	data, err := c.drive.Read(ctx, c.bootSectorAddress, c.bootSectorSize)
	if err != nil {
		return err
	}
	err = c.memory.Load(ctx, c.bootAddress, data)
	if err != nil {
		return err
	}
	err = c.cpu.Jump(ctx, c.bootAddress)
	if err != nil {
		return err
	}
	return c.cpu.Execute(ctx)
}

// NewComputer create new computer
func NewComputer(
	cpu cpu,
	drive drive,
	memory memory,
	bootAddress uint64,
	bootSectorAddress uint64,
	bootSectorSize uint64,
) Computer {
	return &computer{
		bootAddress:       bootAddress,
		bootSectorAddress: bootSectorAddress,
		bootSectorSize:    bootSectorSize,
		cpu:               cpu,
		drive:             drive,
		memory:            memory,
	}
}

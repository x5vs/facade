package main

import (
	"context"
	"fmt"

	"github.com/x5vs/facade/internal/facade"
	"github.com/x5vs/facade/internal/facade/cpu"
	"github.com/x5vs/facade/internal/facade/drive"
	"github.com/x5vs/facade/internal/facade/memory"
)

func main() {
	ctx := context.Background()

	// Create facade for computer №1
	fmt.Println("Create computer №1 facade")
	processor := cpu.NewIntel()
	disk := drive.NewHDD(1024)
	mem := memory.NewDDR1(256)
	computer1 := facade.NewComputer(processor, disk, mem, 0, 0, 128)
	fmt.Println("Start computer №1 facade")
	err := computer1.Start(ctx)
	if err != nil {
		fmt.Println("Computer №1 error: ", err)
	}

	// Create facade for computer №2
	fmt.Println("\n\nCreate computer №2 facade")
	processor = cpu.NewAMD()
	disk = drive.NewSDD(2048)
	mem = memory.NewDDR2(512)
	computer2 := facade.NewComputer(processor, disk, mem, 0, 0, 128)
	fmt.Println("Start computer №2 facade")
	err = computer2.Start(ctx)
	if err != nil {
		fmt.Println("Computer №2 error: ", err)
	}
}

package main

import (
	"fmt"

	"github.com/x5vs/facade/internal/facade/drive"
	"github.com/x5vs/facade/internal/facade/memory"

	"github.com/x5vs/facade/internal/facade"
	"github.com/x5vs/facade/internal/facade/cpu"
)

func case1() facade.Computer {
	processor := cpu.NewIntel()
	disk := drive.NewHDD(1024)
	mem := memory.NewDDR1(256)
	return facade.NewComputer(processor, disk, mem)
}

func case2() facade.Computer {
	processor := cpu.NewAMD()
	disk := drive.NewSDD(2048)
	mem := memory.NewDDR2(512)
	return facade.NewComputer(processor, disk, mem)
}

func main() {
	fmt.Println("Create computer 1 facede")
	computer1 := case1()
	fmt.Println("Start computer 1 facede")
	computer1.Start()
	fmt.Println("\n\nCreate computer 2 facede")
	computer2 := case2()
	fmt.Println("Start computer 2 facede")
	computer2.Start()
}

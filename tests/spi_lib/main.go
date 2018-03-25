package main

import (
	"fmt"

	"golang.org/x/exp/io/spi"
)

const (
	spiDevice = "/dev/spidev0.1"
	spiMode   = spi.Mode0
	spiSpeed  = 10000
)

func main() {
	wb := make([]byte, 4)
	for i, _ := range wb {
		wb[i] = byte((i + 1) * 10)
	}

	rb := make([]byte, 4)

	fmt.Println("Check buffers before spi send:")
	for i := range wb {
		fmt.Println("WB: ", i, "= ", wb[0])
		fmt.Println("RB: ", i, "= ", rb[0])
	}

	var err error
	conn, err := spi.Open(&spi.Devfs{
		Dev:      spiDevice,
		Mode:     spiMode,
		MaxSpeed: spiSpeed,
	})
	if err != nil {
		fmt.Println(err)
	}

	conn.Tx(wb, rb)

	fmt.Println("Check buffers after spi send:")
	for i := range wb {
		fmt.Println("WB: ", i, "= ", wb[0])
		fmt.Println("RB: ", i, "= ", rb[0])
	}
}

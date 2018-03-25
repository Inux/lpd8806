package main

import (
	"fmt"
	"time"

	"github.com/inux/lpd8806"
	"github.com/kidoman/embd"
)

const (
	clkPin  string = "SPI0_SCLK"
	dataPin string = "SPI0_MOSI"
)

func main() {
	var lpd = &lpd8806.LPD8806{}
	lpd.Init(clkPin, 1, dataPin, 2)
	//has to be used in main thread!
	defer embd.CloseGPIO()

	fmt.Println("LED's should be \tWhite")
	lpd.AllToColor(0xFF, 0xFF, 0xFF)
	time.Sleep(time.Second * 2)

	fmt.Println("LED's should be \tBlue")
	lpd.AllToColor(0x00, 0x00, 0xFF)
	time.Sleep(time.Second * 2)

	fmt.Println("LED's should be \tGreen")
	lpd.AllToColor(0x00, 0xFF, 0x00)
	time.Sleep(time.Second * 2)

	fmt.Println("LED's should be \tGreen & Blue")
	lpd.AllToColor(0x00, 0xFF, 0xFF)
	time.Sleep(time.Second * 2)

	fmt.Println("LED's should be \tRed")
	lpd.AllToColor(0xFF, 0x00, 0x00)
	time.Sleep(time.Second * 2)

	fmt.Println("LED's should be \tRed & Blue")
	lpd.AllToColor(0xFF, 0x00, 0xFF)
	time.Sleep(time.Second * 2)

	fmt.Println("LED's should be \tRed & Green")
	lpd.AllToColor(0xFF, 0xFF, 0x00)
	time.Sleep(time.Second * 2)

	fmt.Println("LED's should be \tWhite")
	lpd.AllToColor(0xFF, 0xFF, 0xFF)
	time.Sleep(time.Second * 2)

	fmt.Println("LED's should be \tOff")
	lpd.AllToColor(0x00, 0x00, 0x00)
}

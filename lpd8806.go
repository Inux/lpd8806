package main

import (
	"errors"
	"time"

	"github.com/inux/hwspi"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi" // This loads the RPi driver
)

//LPD8806 is a driver for lpd8806 led strips
type LPD8806 struct {
	LedCount    uint
	colorsCount uint
	latchCount  uint
	ColorArray  []byte
	writeBuffer []byte
	spiCon      *hwspi.HWspi
}

func (lpd *LPD8806) allToColor(r, g, b byte) {
	var i uint
	for i = 0; i < lpd.colorsCount; i = i + 3 {
		lpd.ColorArray[i] = (g | 0x80)
		lpd.ColorArray[i+1] = (b | 0x80)
		lpd.ColorArray[i+2] = (r | 0x80)
	}
	lpd.writeExp()
}

func (lpd *LPD8806) singleToColor(b byte) {

}

func (lpd *LPD8806) segmentToColor(b byte) {

}

func (lpd *LPD8806) allToColorByArray(b byte) {

}

func (lpd *LPD8806) allOff() {
	lpd.initColorArray()
	lpd.writeExp()
}

func (lpd *LPD8806) writeExp() {
	for i := range lpd.writeBuffer {
		lpd.writeBuffer[i] = reverseByte(lpd.ColorArray[i])
	}

	lpd.spiCon.GpioWriteBuffer(lpd.writeBuffer)
}

func (lpd *LPD8806) latch() {
	var i uint
	for i = 0; i < lpd.latchCount; i++ {
		lpd.spiCon.GpioWrite(0x00)
	}
}

func (lpd *LPD8806) initColorArray() {
	var i uint
	count := lpd.colorsCount + lpd.latchCount
	lpd.ColorArray = make([]byte, count)
	lpd.writeBuffer = make([]byte, count)
	for i = 0; i < lpd.colorsCount; i++ {
		lpd.ColorArray[i] = 0x80
	}
	for i = lpd.colorsCount; i < count; i++ {
		lpd.ColorArray[i] = 0x00
	}
}

//Init Creates a New LPD8806 driver
func (lpd *LPD8806) Init(ClkPin, DataPin string, LedCount uint) (*LPD8806, error) {
	if LedCount > 0 {
		lpd.LedCount = LedCount
		lpd.colorsCount = LedCount * 3
		lpd.latchCount = (LedCount + 31) / 32

		lpd.initColorArray()

		lpd.spiCon = &hwspi.HWspi{}
		lpd.spiCon.Init(ClkPin, DataPin, 1)

		//first led commands
		lpd.latch()
		lpd.allOff()
		lpd.allToColor(255, 255, 255)

		return lpd, nil
	}
	return nil, errors.New("LPD8806: Invalid parameters - cannot create LPD8806")
}

const (
	clkPin  string = "SPI0_SCLK"
	dataPin string = "SPI0_MOSI"
)

func reverseByte(b byte) byte {
	r := b
	b >>= 1

	for i := 0; i < 8; i++ {
		r <<= 1
		r |= (byte)(b & 1)
		b >>= 1
	}

	return r
}

func main() {
	var lpd8806 = &LPD8806{}
	lpd8806.Init(clkPin, dataPin, 2)
	//has to be used in main thread!
	defer embd.CloseGPIO()

	time.Sleep(time.Second * 2)
	lpd8806.allToColor(0xFF, 0xFF, 0xFF)
	time.Sleep(time.Second * 2)
	lpd8806.allToColor(0x00, 0x00, 0xFF)
	time.Sleep(time.Second * 2)
	lpd8806.allToColor(0x00, 0xFF, 0x00)
	time.Sleep(time.Second * 2)
	lpd8806.allToColor(0x00, 0xFF, 0xFF)
	time.Sleep(time.Second * 2)
	lpd8806.allToColor(0xFF, 0x00, 0x00)
	time.Sleep(time.Second * 2)
	lpd8806.allToColor(0xFF, 0x00, 0xFF)
	time.Sleep(time.Second * 2)
	lpd8806.allToColor(0xFF, 0xFF, 0x00)
	time.Sleep(time.Second * 2)
	lpd8806.allToColor(0xFF, 0xFF, 0xFF)
	time.Sleep(time.Second * 2)
	lpd8806.allToColor(0x00, 0x00, 0x00)
}

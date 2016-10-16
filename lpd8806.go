package lpd8806

import (
	"errors"
	"time"

	"github.com/inux/hwspi"
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

//AllToColor - all LEDs of the strip to a specific color
func (lpd *LPD8806) AllToColor(r, g, b byte) {
	var i uint
	for i = 0; i < lpd.colorsCount; i = i + 3 {
		lpd.ColorArray[i] = (g | 0x80)
		lpd.ColorArray[i+1] = (r | 0x80)
		lpd.ColorArray[i+2] = (b | 0x80)
	}
	lpd.writeExp()
}

//SingleToColor - single LED of the strip to a specific color
func (lpd *LPD8806) SingleToColor(b byte) {

}

//SegmentToColor - segment of the strip to specific color
func (lpd *LPD8806) SegmentToColor(b byte) {

}

//AllToColorByArray - array will get applied on strip (r,g,b for each LED)
func (lpd *LPD8806) AllToColorByArray(b byte) {

}

//AllOff - switch off the strip (r,g,b to 0)
func (lpd *LPD8806) AllOff() {
	lpd.initColorArray()
	lpd.writeExp()
}

//Init Creates a New LPD8806 driver
func (lpd *LPD8806) Init(ClkPin string, ClkFactor time.Duration,
	DataPin string, LedCount uint) (*LPD8806, error) {
	if LedCount > 0 {
		lpd.LedCount = LedCount
		lpd.colorsCount = LedCount * 3
		lpd.latchCount = (LedCount + 31) / 32

		lpd.initColorArray()

		lpd.spiCon = &hwspi.HWspi{}
		lpd.spiCon.Init(ClkPin, DataPin, ClkFactor)
		//first led commands
		lpd.latch()
		lpd.AllOff()

		return lpd, nil
	}
	return nil, errors.New("LPD8806: Invalid parameters - cannot create LPD8806")
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

func reverseByte(b byte) byte {
	var r = b
	b >>= 1

	for i := 0; i < 8; i++ {
		r <<= 1
		r |= (byte)(b & 1)
		b >>= 1
	}

	return r
}

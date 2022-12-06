package _06

type Device struct {
	input          chan rune
	onStartPacket  func(i int)
	onStartMessage func(i int)
}

func (d *Device) Start() {
	startPacketBuf := make([]rune, 4)
	startMessageBuf := make([]rune, 14)
	var i int
	for r := range d.input {
		i++
		startPacketBuf = append(startPacketBuf[1:4], r)
		startMessageBuf = append(startMessageBuf[1:14], r)
		if isAllDifferent(startPacketBuf) {
			d.onStartPacket(i)
		}
		if isAllDifferent(startMessageBuf) {
			d.onStartMessage(i)
		}
	}
}

func isAllDifferent(buf []rune) bool {
	for i := 0; i < len(buf)-1; i++ {
		if buf[i] == 0 {
			return false
		}
		for j := i + 1; j < len(buf); j++ {
			if buf[i] == buf[j] {
				return false
			}
		}
	}
	return true
}

func NewDevice() *Device {
	return &Device{
		input:          make(chan rune),
		onStartPacket:  func(i int) {},
		onStartMessage: func(i int) {},
	}
}

package bridge

type TV struct {
	enabled bool
	volume  int
	channel int
}

func NewTV() *TV {
	return &TV{
		enabled: false,
		volume:  0,
		channel: 0,
	}
}

func (t *TV) IsEnabled() bool {
	return t.enabled == true
}

func (t *TV) Enable() bool {
	t.enabled = true
	return t.enabled
}

func (t *TV) Disable() bool {
	t.enabled = false
	return true
}

func (t *TV) GetVolume() int {
	return t.volume
}

func (t *TV) SetVolume(volume int) bool {
	t.volume = volume
	return true
}

func (t *TV) GetChannel() int {
	return t.channel
}

func (t *TV) SetChannel(i int) bool {
	t.channel = i
	return true
}

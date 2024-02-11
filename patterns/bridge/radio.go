package bridge

type Radio struct {
	enabled bool
	volume  int
	channel int
}

func (r *Radio) IsEnabled() bool {
	return r.enabled == true
}

func (r *Radio) Enable() bool {
	r.enabled = true
	return true
}

func (r *Radio) Disable() bool {
	r.enabled = false
	return false
}

func (r *Radio) GetVolume() int {
	return r.volume
}

func (r *Radio) SetVolume(volume int) bool {
	r.volume = volume
	return true
}

func (r *Radio) GetChannel() int {
	return r.channel
}

func (r *Radio) SetChannel(i int) bool {
	//TODO implement me
	panic("implement me")
}

func NewRadio() *Radio {
	return &Radio{}
}

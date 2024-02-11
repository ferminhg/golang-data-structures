package bridge

type Device interface {
	IsEnabled() bool
	Enable() bool
	Disable() bool
	GetVolume() int
	SetVolume(int) bool
	GetChannel() int
	SetChannel(int) bool
}

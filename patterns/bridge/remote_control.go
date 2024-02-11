package bridge

type RemoteControl struct {
	device Device
}

func NewRemoteControl(device Device) *RemoteControl {
	return &RemoteControl{device: device}
}

func (r *RemoteControl) TogglePower() bool {
	if r.device.IsEnabled() {
		return r.device.Disable()
	}
	return r.device.Enable()
}

func (r *RemoteControl) VolumeDown() bool {
	return r.device.SetVolume(r.device.GetVolume() - 10)
}

func (r *RemoteControl) VolumeUp() bool {
	return r.device.SetVolume(r.device.GetVolume() + 10)
}

func (r *RemoteControl) ChannelDown() bool {
	return r.device.SetChannel(r.device.GetChannel() - 1)
}

func (r *RemoteControl) ChannelUp() bool {
	return r.device.SetChannel(r.device.GetChannel() + 1)
}

type AdvancedRemoteControl struct {
	RemoteControl
}

func NewAdvancedRemoteControl(device Device) *AdvancedRemoteControl {
	return &AdvancedRemoteControl{RemoteControl{device: device}}
}

func (a *AdvancedRemoteControl) Mute() bool {
	return a.device.SetVolume(0)
}

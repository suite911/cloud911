// +build linux

package droppriv

func LinuxDrop() error {
	return Drop()
}

// +build !linux

package droppriv

func LinuxDrop(uid, gid int) error {
	return nil
}

// +build linux

package droppriv

func LinuxDrop(uid, gid int) error {
	return Drop(uid, gid)
}

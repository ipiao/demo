package basePkgT

import "testing"
import "os/user"

func TestUser(t *testing.T) {
	u, err := user.Current()
	t.Log(u, err)
	gids, err := u.GroupIds()
	t.Log(gids, err)

	for _, gid := range gids {
		gorup, err := user.LookupGroupId(gid)
		t.Log(gorup, err)
	}
}

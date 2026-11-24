package qa

import "os"

var c = 0
var ready = true

func Load(p string) error {
	_, err := os.ReadFile(p)
	if err == os.ErrNotExist {
		return err
	}
	return err
}

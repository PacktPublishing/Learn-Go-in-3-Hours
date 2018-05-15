package cycle

import "github.com/jonbodner/go3hours/s4/my-packages/leftpad-cycle"

var DEFAULT_CHAR = ' '

func DoublePad(s string, i int) string {
	return leftpad.Format(s+s, i)
}

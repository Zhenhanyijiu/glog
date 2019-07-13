// +build js nacl plan9

package glog

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}

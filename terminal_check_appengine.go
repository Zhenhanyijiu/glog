// +build appengine

package glog

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return true
}

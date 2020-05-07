// +build !windows

package lcf

import (
	"errors"
)

func windowsNativeANSI(_ bool, _ bool, _ interface{}) (bool, error) {
	return false, errors.New("not available on this platform")
}

//go:build !plan9 && !windows
// +build !plan9,!windows

package gitpkg

import "fmt"

// preReceiveHook returns the bytes of a pre-receive hook script
// that prints m before exiting successfully
func preReceiveHook(m string) []byte {
	return []byte(fmt.Sprintf("#!/bin/sh\nprintf '%s'\n", m))
}

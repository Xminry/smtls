//go:build js
// +build js

package smtls

var (
	hasGCMAsmAMD64 = false
	hasGCMAsmARM64 = false
	hasGCMAsmS390X = false

	hasAESGCMHardwareSupport = false
)

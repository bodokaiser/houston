// Package driver provides sysfs implementation of device drivers.
//
// The subpackages are structured by device function and device family.
package driver

// Driver implements a device driver.
type Driver interface {
	Init() error
}

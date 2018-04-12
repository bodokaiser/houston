// Package model provides entities for safe device configuration.
//
// By using an extra model layer between the device drivers and network
// interfaces like HTTP we can control what information we want to expose
// to the outside world. On the one hand side we can catch invalid device
// configurations at an early stage on the other side we can hide
// (complicated) implementation details from the API user.
package model

//go:build !linux
// +build !linux

package osarch

func ArchitectureGetLocal() (string, error) {
	return ArchitectureDefault, nil
}

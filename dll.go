package gonfapi

import "golang.org/x/sys/windows"

type SafeDll struct {
	dll *windows.LazyDLL
}

func (sd *SafeDll) Clean() {
	sd.dll = nil
}
func (sd *SafeDll) Loading(dll string) error {
	sd.dll = windows.NewLazyDLL(dll)
	return sd.dll.Load()
}

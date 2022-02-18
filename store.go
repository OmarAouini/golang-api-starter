package main

type Store interface {
	GetRestreamers() ([]Restreamer, error)
}

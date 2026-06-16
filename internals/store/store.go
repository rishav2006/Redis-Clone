package store

import "sync"

type Store struct {
	Data map[string]string
	Mu   sync.RWMutex
}

var DB = Store{
	Data: make(map[string]string),
}
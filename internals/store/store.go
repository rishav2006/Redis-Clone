package store

import (
	"sync"
	"time"
)

type Store struct {
	Data       map[string]string
	Mu         sync.RWMutex
	Expiration map[string]time.Time
}

var DB = Store{
	Data:       make(map[string]string),
	Expiration: make(map[string]time.Time),
}

// Expiration map[string]time.Time

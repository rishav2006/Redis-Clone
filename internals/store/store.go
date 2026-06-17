package store

import (
	"net"
	"sync"
	"time"
)

type Store struct {
	Data        map[string]string
	Mu          sync.RWMutex
	Expiration  map[string]time.Time
	Subscribers map[string][]net.Conn
}

var DB = Store{
	Data:        make(map[string]string),
	Expiration:  make(map[string]time.Time),
	Subscribers: make(map[string][]net.Conn),
}

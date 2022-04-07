package services

import "time"

type cache struct {
	_m map[string]interface{}
}

// ttls is time to live in seconds
func (s cache) Set(key string, value interface{}, ttls uint) {
	s._m[key] = value
	go func() {
		time.Sleep(time.Second * time.Duration(ttls))
		s.Clear(key)
	}()
}

func (s cache) Get(key string) interface{} {
	return s._m[key]
}

// clear all
func (s cache) ClearAll() {
	for k := range s._m {
		delete(s._m, k)
	}
}

func (s cache) Clear(key string) {
	delete(s._m, key)
}

package services

type cache struct {
	_m map[string]interface{}
}

func (s cache) Set(key string, value interface{}) {
	s._m[key] = value
}

func (s cache) Get(key string) interface{} {
	return s._m[key]
}

func (s cache) Clear() {
	for k := range s._m {
		s._m[k] = nil
	}
}

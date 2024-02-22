package patterns

import "sync"

type singleton struct{}

var ins *singleton = &singleton{}

func GetInsOr() *singleton {
	return ins
}

var ins2 *singleton

func GetInsOr2() *singleton {
	if ins2 == nil {
		ins2 = &singleton{}
	}
	return ins2
}

var mu sync.Mutex

func GetInsSaty() *singleton {
	if ins2 == nil {
		mu.Lock()
		if ins2 == nil {
			ins2 = &singleton{}
		}
		mu.Unlock()
	}
	return ins2
}

var ins3 *singleton
var once sync.Once

func GetInsSafy2() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}

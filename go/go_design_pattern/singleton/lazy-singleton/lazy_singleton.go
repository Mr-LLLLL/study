package lazy_singleton

import "sync"

var (
	lazySingleton *singleton
	once          = &sync.Once{}
)

type singleton struct{}

func GetLazySingleton() *singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &singleton{}
		})
	}
	return lazySingleton
}

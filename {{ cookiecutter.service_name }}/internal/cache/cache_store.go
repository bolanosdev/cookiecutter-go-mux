package cache

type InMemoryCacheStore struct {
}

var instantiated *InMemoryCacheStore = nil

func NewCacheStore() *InMemoryCacheStore {
	cache_store := InMemoryCacheStore{}
	return &cache_store
}

func GetCacheStore() *InMemoryCacheStore {
	if instantiated == nil {
		instantiated = NewCacheStore()
	}
	return instantiated
}

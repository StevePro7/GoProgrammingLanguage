package memo

import "sync"

// Func is the type of the function to memoize
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]result
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// Get is concurrency safe
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	defer memo.mu.Unlock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)

		// Between the two critical sections, several goroutines
		// may race to compute f(key) and update the map
		memo.mu.Lock()
		defer memo.mu.Unlock()
		memo.cache[key] = res
	}
	return res.value, res.err
}

// Package memo provides a concurrency-safe memoization a function of
// type Func. Concurrent requests are serialized by a Mutex.
package memo

// Func is the type of the function to memoize.
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

// A request is a message requesting that the Func be applied to the key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
}

type Memo struct {
	requests chan request
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}
func (memo *Memo) Get(key string) (value interface{}, err error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close()  {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate teh function.
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for ready condition.
	<-e.ready
	// Send the result to the client.
	response <- e.res
}

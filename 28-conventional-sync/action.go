package main

import (
	"context"
	"sync"
	"time"
)

// Embed mutexes into other types to make their operations safe.
// Performance overhead.
type SafeMap struct {
	sync.Mutex
	m map[string]int
}

func (sm *SafeMap) Incr(key string) {
	sm.Lock()
	defer sm.Unlock() // <- Good habit to defer. Avoid mistakes.

	sm.m[key]++

}

// Sometimes we need to prefer readers to (infrequent) writers
// 99% chance is I will read and not write

type InfoClient struct {
	mu        sync.RWMutex
	token     string
	tokenTime time.Time
	TTL       time.Duration
}

func (i *InfoClient) CheckToken() (string, time.Duration) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	return i.token, i.TTL - time.Since(i.tokenTime)
}

func (i *InfoClient) ReplaceToken(ctx context.Context) (string, error) {
	// Do the expensive part first before lock/unlock.
	token, ttl, err := i.getAccesToken(ctx)

	if err != nil {
		return "", err
	}

	i.mu.Lock() // Must replace tocken, no other coroutine should be able to access it.
	defer i.mu.Unlock()

	i.token = token
	i.tokenTime = time.Now()
	i.TTL = time.Duration(ttl) * time.Second

	return token, nil

}

// empty function, would have some logic
func (i *InfoClient) getAccesToken(ctx context.Context) (string, time.Duration, error) {
	return "", 1 * time.Second, nil
}

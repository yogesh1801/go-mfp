// MFP - Miulti-Function Printers and scanners toolkit
// DNS-SD service discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Generic sets

package generic

import "sync"

// LockedSet is the generic set of any comparable objects.
// It works like [Set], but goroutine-safe.
type LockedSet[T comparable] struct {
	set  Set[T]     // Underlying Set
	lock sync.Mutex // Access lock
}

// NewLockedSet creates a new LockedSet
func NewLockedSet[T comparable]() *LockedSet[T] {
	return &LockedSet[T]{set: NewSet[T]()}
}

// Clear purges the set
func (ls *LockedSet[T]) Clear() {
	ls.lock.Lock()
	ls.set.Clear()
	ls.lock.Unlock()
}

// Empty reports if set is empty
func (ls *LockedSet[T]) Empty() bool {
	ls.lock.Lock()
	defer ls.lock.Unlock()
	return ls.set.Empty()
}

// Contains reports if member already in the set
func (ls *LockedSet[T]) Contains(member T) bool {
	ls.lock.Lock()
	defer ls.lock.Unlock()
	return ls.set.Contains(member)
}

// Add adds member to the set
func (ls *LockedSet[T]) Add(member T) {
	ls.lock.Lock()
	ls.set.Add(member)
	ls.lock.Unlock()
}

// Del deletes member from the set
func (ls *LockedSet[T]) Del(member T) {
	ls.lock.Lock()
	ls.set.Del(member)
	ls.lock.Unlock()
}

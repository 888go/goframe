// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package mutex provides switch of concurrent safe feature for sync.Mutex.
package mutex

import (
	"sync"
)

// Mutex is a sync.Mutex with a switch for concurrent safe feature.
type Mutex struct {
	// Underlying mutex.
	mutex *sync.Mutex
}

// New creates and returns a new *Mutex.
// The parameter `safe` is used to specify whether using this mutex in concurrent safety,
// which is false in default.

// ff:
// safe:
func New(safe ...bool) *Mutex {
	mu := Create(safe...)
	return &mu
}

// Create creates and returns a new Mutex object.
// The parameter `safe` is used to specify whether using this mutex in concurrent safety,
// which is false in default.

// ff:
// safe:
func Create(safe ...bool) Mutex {
	if len(safe) > 0 && safe[0] {
		return Mutex{
			mutex: new(sync.Mutex),
		}
	}
	return Mutex{}
}

// IsSafe checks and returns whether current mutex is in concurrent-safe usage.

// ff:
func (mu *Mutex) IsSafe() bool {
	return mu.mutex != nil
}

// Lock locks mutex for writing.
// It does nothing if it is not in concurrent-safe usage.

// ff:
func (mu *Mutex) Lock() {
	if mu.mutex != nil {
		mu.mutex.Lock()
	}
}

// Unlock unlocks mutex for writing.
// It does nothing if it is not in concurrent-safe usage.

// ff:
func (mu *Mutex) Unlock() {
	if mu.mutex != nil {
		mu.mutex.Unlock()
	}
}

package main

import "sync"

var bufferPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 4096) // Create a new 4KB buffer
	},
}

func Process() {
	buf := bufferPool.Get().([]byte) // Get a buffer
	defer bufferPool.Put(buf)        // Put it back when done

	// ... use buf for I/O ...
}

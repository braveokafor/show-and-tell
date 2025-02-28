package handlers

import (
	"fmt"
	"net/http"
	"sync"
)

// Global variable to hold allocated memory to prevent garbage collection
var (
	memoryHog    [][]byte
	memoryMutex  sync.Mutex
	totalAllocMB int
)

// allocateMemory allocates memory and keeps it allocated (to demonstrate auto-scaling)
func (h *Handler) allocateMemory(w http.ResponseWriter, r *http.Request) {
	memoryMutex.Lock()
	defer memoryMutex.Unlock()

	// Allocate 10MB of memory each time
	chunkSize := 10 * 1024 * 1024 // 10MB

	// Create a new chunk and fill it with data
	newChunk := make([]byte, chunkSize)
	for i := 0; i < chunkSize; i++ {
		newChunk[i] = byte(i % 256)
	}

	// Add to our slice of memory chunks
	memoryHog = append(memoryHog, newChunk)
	totalAllocMB += 10

	// Return status
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Total memory allocated: %d MB\n", totalAllocMB)
	fmt.Fprintf(w, "Number of chunks: %d\n", len(memoryHog))
	fmt.Fprintf(w, "Refresh this page to allocate more memory\n")
}

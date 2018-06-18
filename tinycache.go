package tinycache

import (
	"sync"
)

const defaultMaxElements = 100

type Cache struct {
	id          int
	maxElements int
	mu          sync.RWMutex
	elements    []interface{}
}

// NewCache returns a new history object
// with the default number of elements
func NewCache() *Cache {
	return &Cache{
		maxElements: defaultMaxElements,
	}
}

func (h *Cache) SetMaxElements(i int) {
	h.maxElements = i
}

func (h *Cache) Add(newElement interface{}) {
	h.mu.Lock()

	if len(h.elements) >= h.maxElements {
		copy(h.elements, h.elements[len(h.elements)-h.maxElements+1:])
		h.elements = h.elements[:h.maxElements-1]
	}
	h.elements = append(h.elements, newElement)

	h.mu.Unlock()
}

func (h *Cache) GetElements() []interface{} {
	return h.elements
}

func (h *Cache) Count() int {
	return len(h.elements)
}

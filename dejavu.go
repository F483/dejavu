package dejavu

import (
	"crypto/sha256"
	"sync"
)

type digest [sha256.Size]byte

type DejaVu struct {
	buffer []digest       // ring buffer
	size   int            // ring buffer size
	index  int            // current ring buffer index
	lookup map[digest]int // digest -> newest index (for performance)
	mutex  *sync.Mutex
}

// Creates a new DejaVu memory with max entries limited to given size.
func NewDejaVu(size int) *DejaVu {
	// FIXME handle size of < 1 given
	return &DejaVu{
		buffer: make([]digest, size),
		size:   size,
		index:  0,
		lookup: make(map[digest]int),
		mutex:  new(sync.Mutex),
	}
}

// Same as Witness method but bypasses hashing the data. Use this to
// to improve performance if you already happen to have the sha256 digest
// of the data entry.
func (d *DejaVu) WitnessDigest(dataDigest [sha256.Size]byte) bool {
	d.mutex.Lock()
	_, familiar := d.lookup[dataDigest] // check if previously seen

	// rm oldest lookup key if no newer entry
	maxed := len(d.buffer) == d.size // overwriting oldest entry
	if maxed && (d.lookup[d.buffer[d.index]] == d.index) {
		delete(d.lookup, d.buffer[d.index]) // no newer entries
	}

	// add entry and update index/lookup
	d.buffer[d.index] = dataDigest
	d.lookup[dataDigest] = d.index
	d.index = (d.index + 1) % d.size

	d.mutex.Unlock()
	return familiar
}

// Add data entry to memory. Returns true if previously seen, may give
// false negatives but not false positives.
func (d *DejaVu) Witness(data []byte) bool {
	return d.WitnessDigest(sha256.Sum256(data))
}

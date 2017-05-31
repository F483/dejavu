package dejavu

import (
	"crypto/sha256"
	"github.com/AndreasBriese/bbloom"
	"sync"
)

// DejaVu witnesses data and recalls if seen before.
type DejaVu interface {

	// Witness data and add to memory. Returns true if previously seen.
	Witness(data []byte) bool

	// WitnessDigest is equivalent to the Winness method but bypasses hashing
	// the data. Use this to improve performance if you already happen
	// to have the sha256 digest.
	WitnessDigest(digest [sha256.Size]byte) bool
}

//////////////////////////////////
// DETERMINISTIC IMPLEMENTATION //
//////////////////////////////////

type deterministic struct {
	buffer [][sha256.Size]byte       // ring buffer
	size   int                       // ring buffer size
	index  int                       // current ring buffer index
	lookup map[[sha256.Size]byte]int // digest -> newest index (optimization)
	mutex  *sync.Mutex
}

// NewDeterministic creates a deterministic DejaVu memory. Will remember
// most recent entries within given entrie limit and forget older entries.
func NewDeterministic(entrieLimit uint) DejaVu {
	return &deterministic{
		buffer: make([][sha256.Size]byte, entrieLimit),
		size:   int(entrieLimit),
		index:  0,
		lookup: make(map[[sha256.Size]byte]int),
		mutex:  new(sync.Mutex),
	}
}

func (d *deterministic) WitnessDigest(digest [sha256.Size]byte) bool {
	d.mutex.Lock()

	_, familiar := d.lookup[digest] // check if previously seen

	// rm oldest lookup key if no newer entry
	maxed := len(d.buffer) == d.size // overwriting oldest entry
	if maxed && (d.lookup[d.buffer[d.index]] == d.index) {
		delete(d.lookup, d.buffer[d.index]) // no newer entries
	}

	// add entry and update index/lookup
	d.buffer[d.index] = digest
	d.lookup[digest] = d.index
	d.index = (d.index + 1) % d.size

	d.mutex.Unlock()
	return familiar
}

func (d *deterministic) Witness(data []byte) bool {
	return d.WitnessDigest(sha256.Sum256(data))
}

//////////////////////////////////
// PROBABILISTIC IMPLEMENTATION //
//////////////////////////////////

type probabilistic struct {
	filters            [2]*bbloom.Bloom // alternatingly replaced every size entries
	entrieLimit        uint             // filter size
	falsePositiveRatio float64          //
	index              int              // current filter index
	entries            uint             // entries added to currently indexed filter
	mutex              *sync.Mutex
}

// NewProbabilistic creates a probabilistic DejaVu memory. Probably remembers
// most recent entries within given entrie limit and false positive ratio.
func NewProbabilistic(entrieLimit uint, falsePositiveRatio float64) DejaVu {
	a := bbloom.New(float64(entrieLimit), falsePositiveRatio)
	b := bbloom.New(float64(entrieLimit), falsePositiveRatio)
	return &probabilistic{
		filters:            [2]*bbloom.Bloom{&a, &b},
		entrieLimit:        entrieLimit,
		falsePositiveRatio: falsePositiveRatio,
		index:              0,
		entries:            0,
		mutex:              new(sync.Mutex),
	}
}

func (p *probabilistic) WitnessDigest(digest [sha256.Size]byte) bool {
	p.mutex.Lock()

	// check if exists
	d := digest[:]
	familiar := p.filters[0].Has(d) || p.filters[1].Has(d)

	// always add in case its from the old buffer
	p.filters[p.index].AddIfNotHas(d)
	p.entries++

	// switch buffers if current is maxed
	if p.entries >= p.entrieLimit {
		p.entries = 0
		p.index = (p.index + 1) % 2
		f := bbloom.New(float64(p.entrieLimit), p.falsePositiveRatio)
		p.filters[p.index] = &f // replace old filter
	}

	p.mutex.Unlock()
	return familiar
}

func (p *probabilistic) Witness(data []byte) bool {
	return p.WitnessDigest(sha256.Sum256(data))
}

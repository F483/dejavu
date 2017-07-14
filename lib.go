/*
Package dejavu offers quick detection of already witnessed data.

Limited memory of witnessed data, oldest are forgotten. Library is
thread safe. Offers deterministic and probabilistic (over an order of
magnatude less memory consuming) implementation.
*/
package dejavu

import (
	"bufio"
	"crypto/sha256"
	"github.com/willf/bloom"
	"io"
	"os"
	"sync"
)

// Version information
const Version string = "0.1.0"

// DejaVu witnesses data and recalls if seen before.
type DejaVu interface {

	// Witness data and add to memory. Returns true if previously seen.
	Witness(data []byte) bool

	// WitnessDigest is equivalent to the Winness method but bypasses
	// hashing the data. Use this to improve performance if you already
	// happen to have the sha256 digest.
	WitnessDigest(digest [sha256.Size]byte) bool
}

// New creates a probabilistic or deterministic DejaVu memory with given
// entrie limit and false positive ratio (only used for probabilistic).
func New(probabilistic bool, limit uint32, fpRatio float64) DejaVu {
	if probabilistic {
		return NewProbabilistic(limit, fpRatio)
	}
	return NewDeterministic(limit)
}

///////////////////////////////////
// PROCESS TEXT (for dejavu bin) //
///////////////////////////////////

func getReaders(paths []string) []io.Reader {
	readers := make([]io.Reader, len(paths))
	for i, path := range paths {
		if path == "-" { // read from stdin
			readers[i] = os.Stdin
		} else { // read from file path
			file, err := os.Open(path)
			if err != nil {
				panic(err) // TODO bettor error handling?
			}
			readers[i] = file
		}
	}
	return readers
}

func getWriter(path string) io.Writer {
	if path == "" {
		return os.Stdout
	}
	file, err := os.Open(path)
	if err != nil {
		panic(err) // TODO bettor error handling?
	}
	return file
}

// ProcessPaths is equivalent to Process, only that file paths are given.
// If - in inputs to use stdin and empty out to use stdout.
func ProcessPaths(d DejaVu, filter bool, out string, inputs ...string) {
	writer := getWriter(out)
	readers := getReaders(inputs)
	Process(d, filter, writer, readers...)
}

// Process given inputs as text to output with dejavu instance.
// If filter is true duplicates are filtered, otherwise only
// duplicates sent to output.
func Process(d DejaVu, filter bool, out io.Writer, inputs ...io.Reader) {
	for _, input := range inputs {
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			text := scanner.Text()
			seen := d.Witness([]byte(text))
			if (!filter && seen) || (filter && !seen) {
				out.Write([]byte(text))
			}
		}
	}
}

//////////////////////////////////
// DETERMINISTIC IMPLEMENTATION //
//////////////////////////////////

type deterministic struct {
	buffer [][sha256.Size]byte       // ring buffer
	size   int                       // ring buffer size
	index  int                       // current ring buffer index
	lookup map[[sha256.Size]byte]int // digest -> newest index
	mutex  *sync.Mutex
}

// NewDeterministic creates a deterministic DejaVu memory. Will remember
// most recent entries within given entrie limit and forget older entries.
func NewDeterministic(limit uint32) DejaVu {
	return &deterministic{
		buffer: make([][sha256.Size]byte, limit),
		size:   int(limit),
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

const liveFilterCnt = 8
const totalFilterCnt = liveFilterCnt + 1

type probabilistic struct {
	filters            []*bloom.BloomFilter
	limit              uint32  // filter size
	falsePositiveRatio float64 // remember for buffer switch
	index              int     // current filter index
	entries            uint32  // entries in currently indexed filter
	mutex              *sync.Mutex
}

// NewProbabilistic creates a probabilistic DejaVu memory. Probably
// remembers most recent entries within given entrie limit and false
// positive ratio. False positive ratio should be between 0.0 and 1.0.
func NewProbabilistic(limit uint32, falsePositiveRatio float64) DejaVu {
	filters := make([]*bloom.BloomFilter, totalFilterCnt, totalFilterCnt)
	for i := 0; i < totalFilterCnt; i++ {
		fl := uint(limit / liveFilterCnt)
		filters[i] = bloom.NewWithEstimates(fl, falsePositiveRatio)
	}
	return &probabilistic{
		filters:            filters,
		limit:              limit,
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
	familiar := false
	for _, f := range p.filters {
		if f.Test(d) {
			familiar = true
			break
		}
	}

	// always add in case its from the old buffer
	p.filters[p.index].Add(d)
	p.entries++

	// switch buffers if current is maxed
	if p.entries >= (p.limit / liveFilterCnt) {
		p.entries = 0
		p.index = (p.index + 1) % len(p.filters)
		fl := uint(p.limit / liveFilterCnt)
		f := bloom.NewWithEstimates(fl, p.falsePositiveRatio)
		p.filters[p.index] = f // replace old filter
	}

	p.mutex.Unlock()
	return familiar
}

func (p *probabilistic) Witness(data []byte) bool {
	return p.WitnessDigest(sha256.Sum256(data))
}

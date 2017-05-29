package dejavu

import "testing"

func TestDejaVu(t *testing.T) {

	d := NewDejaVu(3)

	// add entries
	if d.Witness([]byte("foo")) {
		t.Errorf("Incorrect déjà vu: 'foo'!")
	}
	if d.Witness([]byte("bar")) {
		t.Errorf("Incorrect déjà vu: 'bar'!")
	}

	// remembers entry
	if !d.Witness([]byte("bar")) {
		t.Errorf("Expected déjà vu: 'bar'!")
	}

	// remembers oldest entry before overwriting
	if !d.Witness([]byte("foo")) {
		t.Errorf("Expected déjà vu: 'foo'!")
	}

	// add entries
	if d.Witness([]byte("bam")) {
		t.Errorf("Incorrect déjà vu: 'bam'!")
	}
	if d.Witness([]byte("baz")) {
		t.Errorf("Incorrect déjà vu: 'baz'!")
	}

	// forgot oldest
	if d.Witness([]byte("bar")) {
		t.Errorf("Incorrect déjà vu: 'bar'!")
	}
}

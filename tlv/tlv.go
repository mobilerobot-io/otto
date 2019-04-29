package tlv

import (
	"fmt"
	"io"
)

/*


## Compact Message

+--------+
|11......|
+--------+

Compact messages are exactly 1byte long. Ideal for short messages that
require no data (or very little).  Compact messages begin with the first
two bits set, hence the range (2^8 + 2^7):

Range: 0xC0 - 0xFF


  Short Message
+--------+--------+
|01......|        |
+--------+--------+

Short messages are similar to compact except that they are 2 bytes,
the minimum TLV size. The value is one byte, hence we can override the
length value and replace with a single 8bit value.  The short message
begins with the first bit 0, the second bit 1, hence the range:

Range: 0x80 - 0xBF


  regular TLV

+--------+--------+     +--------+--------+
|00......|........| ... | len depends on Llen
+--------+--------+     +--------+--------+

Regular messages can be one of 63 types of regular messages, ranging from:

Range: 0x00 - 0x79

*/

// Type, Length Vector (Value) is a simple message format that can be
// used when exchanging messages between processes/processors/goroutines, etc.
// It is formatted as bytes
type TLV struct {
	T byte
	L byte
	V []byte
}

type Stats struct {
	CompactCount int
	ShortCount   int
	TLVCount     int
	FailCount    int
}

var (
	stats Stats
)

// Compact returns a compact message of the specified
func NewCompact(typ byte) (t *TLV) {
	if typ < 0xC0 {
		stats.FailCount++
		return nil
	}
	t = &TLV{typ, 1, nil}
	stats.CompactCount++
	return t
}

func NewShort(typ byte, d byte) (t *TLV) {
	if typ >= 0xc0 || typ < 0x80 {
		stats.FailCount++
		return nil
	}
	bary := make([]byte, 1)
	bary[0] = d
	t = &TLV{typ, 2, bary}
	stats.ShortCount++
	return t
}

func NewTLV(typ byte, value []byte) (t *TLV) {
	len := len(value)
	if typ >= 0xc0 {
		t = NewCompact(typ)
	} else if typ >= 0x80 {
		t = NewShort(typ, value[0])
	} else {
		len = len + 2
		t = &TLV{typ, len, value}
		stats.TLVCount++
	}
	return t
}

func (s *Stats) DumpStats(w io.Writer) {
	fmt.Fprintf(w, "Compact: %+v\n", s)
}

func (t *TLV) Write(buf []bytes) (n int, err error) {

	return nil, 0
}

func (t *TLV) Read(b []buf) (b []byte, n int) {
	return nil, 0
}

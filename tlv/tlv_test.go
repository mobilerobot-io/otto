package tlv

import (
	"testing"
)

func TestCompact(t *testing.T) {
	var tlv *TLV
	tsts := []struct {
		T byte
		E bool
	}{
		{0x1, false},
		{0x79, false},
		{0x80, false},
		{0xaa, false},
		{0xbf, false},
		{0xc0, true},
		{0xc1, true},
		{0xd2, true},
		{0xff, true},
	}

	for _, tst := range tsts {
		tlv = NewCompact(tst.T)
		notnil := (tlv != nil)
		if notnil != tst.E {
			t.Errorf("typ (%#02x) expected (%t) got (%t) ", tst.T, tst.E, notnil)
		}
	}
}

func TestShort(t *testing.T) {
	var tlv *TLV
	tsts := []struct {
		T byte
		D byte
		E bool
	}{
		{0x1, 0xbe, false},
		{0x79, 0xff, false},
		{0x80, 0x12, true},
		{0xaa, 0x93, true},
		{0xbf, 0x22, true},
		{0xc0, 0x21, false},
		{0xc1, 0xf8, false},
		{0xd2, 0x21, false},
		{0xff, 0x23, false},
	}

	for _, tst := range tsts {
		tlv = NewShort(tst.T, tst.D)
		notnil := (tlv != nil)
		if notnil != tst.E {
			t.Errorf("typ (%#02x) expected (%t) got (%t) ", tst.T, tst.E, notnil)
		}
		if tlv == nil {
			continue
		}
		if tlv.T != tst.T {
			t.Errorf("test type is different than tlv type")
		}

		if tlv.V[0] != tst.D {
			t.Error("test data tpe is different than tlv")
		}
	}
}

/*
func TestTLV(t *testing.T) {
	var tlv *TLV
	tsts := []struct {
		T byte
		L byte
		D []byte
		E bool
	}{
		{0x2, 0x4, [4]byte{0xbe, 0xef, 0xca, 0xfe}, true},
		{0x4, 0x2, [4]byte{0x0}, true},
		{0x5, 0x3, [4]byte{0xde, 0xed}, true},
		{0x79, 0x4, [4]byte{0xca, 0xfe, 0xde, 0xad}, true},
		{0x80, 0x2, [4]byte{0x2}, true},
		{0xaa, 0x2, [4]byte{0x2}, true},
		{0xbf, 0x2, [4]byte{0x2}, true},
		{0xc0, 0x1, [4]byte{0x21}, true},
		{0xc1, 0x1, [4]byte{0xf8}, true},
		{0xd2, 0x1, [4]byte{0x21}, true},
		{0xff, 0x1, [4]byte{0x23}, true},
	}

	for _, tst := range tsts {
		tlv = NewTLV(tst.T, []byte(tst.D))
		notnil := (tlv != nil)
		if notnil != tst.E {
			t.Errorf("typ (%#02x) expected (%t) got (%t) ", tst.T, tst.E, notnil)
		}
		if tlv == nil {
			continue
		}
		if tlv.T != tst.T {
			t.Errorf("test type is different than tlv type")
		}
		if tlv.L != tst.L {
			t.Errorf("test type expect len (%d) got (%d)", tst.L, tlv.L)
		}

			var d []byte
			if tlv.V != d {
				t.Error("test data tpe is different than tlv")
			}
	}
}
*/

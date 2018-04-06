package ip2long

import (
	"testing"
)

type ipv42longTest struct {
	desc string
	in   string
	out  int
	err  error
}

func (ct *ipv42longTest) check(t *testing.T, i int) {
	out, err := IPv42long(ct.in)
	pass := true
	if err != ct.err {
		pass = false
		t.Errorf("%d:%s:error: got %v; want %v", i, ct.desc, err, ct.err)
	}
	if out != ct.out {
		pass = false
		t.Errorf("%d:%s:out: got %q; want %q", i, ct.desc, out, ct.out)
	}

	if pass {
		t.Logf("%d:%s passed", i, ct.desc)
	}
}

func TestIPv42long(t *testing.T) {
	for i, ct := range []ipv42longTest{{
		desc: "normal 1",
		in:   "172.168.5.1",
		out:  2896692481,
		err:  nil,
	}, {
		desc: "normal 2",
		in:   "0000.0.0.0",
		out:  0,
		err:  nil,
	}, {
		desc: "normal 3",
		in:   "255.255.255.255",
		out:  4294967295,
		err:  nil,
	}, {
		desc: "prefix zeros",
		in:   "172.168.05.0001",
		out:  2896692481,
		err:  nil,
	}, {
		desc: "prefix spaces",
		in:   "    172.  168.  5 .1",
		out:  2896692481,
		err:  nil,
	}, {
		desc: "suffix spaces",
		in:   "172 .168    .5.1  ",
		out:  2896692481,
		err:  nil,
	}, {
		desc: "combination 1",
		in:   " 172   .168  .   05.1",
		out:  2896692481,
		err:  nil,
	}, {
		desc: "spaces between digits",
		in:   "172.16 8.5.1",
		out:  0,
		err:  ErrInvalidIPv4Address,
	}, {
		desc: "overflowed segments",
		in:   "172.568.5.1",
		out:  0,
		err:  ErrOverflowedIPv4Segment,
	}, {
		desc: "invalid characters",
		in:   "172.168.0x05.1",
		out:  0,
		err:  ErrMalformedIPv4Address,
	}, {
		desc: "too many segs",
		in:   "172.168.1.1.1.1",
		out:  0,
		err:  ErrInvalidIPv4Address,
	}, {
		desc: "dots",
		in:   "172....168..5.1",
		out:  0,
		err:  ErrInvalidIPv4Address,
	}} {
		ct.check(t, i)
	}
}

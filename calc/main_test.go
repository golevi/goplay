package main

import "testing"

func TestXor(t *testing.T) {
	if xor(false, false) != false {
		t.Error("xor: false, false, should equal false")
	}

	if xor(false, true) != true {
		t.Error("xor: false, true, should equal true")
	}

	if xor(true, false) != true {
		t.Error("xor: true, false, should equal true")
	}

	if xor(true, true) != false {
		t.Error("xor: true, true, should equal true")
	}
}

func TestAnd(t *testing.T) {
	if and(false, false) != false {
		t.Error("and: false, false, should equal false")
	}

	if and(false, true) != false {
		t.Error("and: false, true, should equal false")
	}

	if and(true, false) != false {
		t.Error("and: true, false, should equal false")
	}

	if and(true, true) != true {
		t.Error("and: true, true, should equal true")
	}
}

func TestOr(t *testing.T) {
	if or(false, false) != false {
		t.Error("or: false, false, should equal false")
	}

	if or(false, true) != true {
		t.Error("or: false, true, should equal true")
	}

	if or(true, false) != true {
		t.Error("or: true, false, should equal true")
	}

	if or(true, true) != true {
		t.Error("or: true, true, should equal true")
	}
}

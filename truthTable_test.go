package main

import "testing"

func TestNegation(t *testing.T) {
	cases := []struct {
		input bool
		want  bool
	}{
		{false, true},
		{true, false},
	}

	for _, c := range cases {
		got := negation(c.input)
		if got != c.want {
			t.Errorf("negation(%t) == %t, want %t", c.input, got, c.want)
		}
	}

}

func TestConjunction(t *testing.T) {
	cases := []struct {
		inputP, inputQ bool
		want           bool
	}{
		{true, true, true},
		{true, false, false},
		{false, true, false},
		{false, false, false},
	}

	for _, c := range cases {
		got := conjunction(c.inputP, c.inputQ)
		if got != c.want {
			t.Errorf("conjunction(%t,%t) == %t, want %t", c.inputP, c.inputQ, got, c.want)
		}
	}

}

func TestInclusiveDisjunction(t *testing.T) {
	cases := []struct {
		inputP, inputQ bool
		want           bool
	}{
		{true, true, true},
		{true, false, true},
		{false, true, true},
		{false, false, false},
	}

	for _, c := range cases {
		got := inclusiveDisjunction(c.inputP, c.inputQ)
		if got != c.want {
			t.Errorf("inclusiveDisjunction(%t,%t) == %t, want %t", c.inputP, c.inputQ, got, c.want)
		}
	}

}

func TestExclusiveDisjunction(t *testing.T) {
	cases := []struct {
		inputP, inputQ bool
		want           bool
	}{
		{true, true, false},
		{true, false, true},
		{false, true, true},
		{false, false, false},
	}

	for _, c := range cases {
		got := exclusiveDisjunction(c.inputP, c.inputQ)
		if got != c.want {
			t.Errorf("exclusiveDisjunction(%t,%t) == %t, want %t", c.inputP, c.inputQ, got, c.want)
		}
	}

}

func TestConditional(t *testing.T) {
	cases := []struct {
		inputP, inputQ bool
		want           bool
	}{
		{true, true, true},
		{true, false, false},
		{false, true, true},
		{false, false, true},
	}

	for _, c := range cases {
		got := conditional(c.inputP, c.inputQ)
		if got != c.want {
			t.Errorf("conditional(%t,%t) == %t, want %t", c.inputP, c.inputQ, got, c.want)
		}
	}

}

func TestBiConditional(t *testing.T) {
	cases := []struct {
		inputP, inputQ bool
		want           bool
	}{
		{true, true, true},
		{true, false, false},
		{false, true, false},
		{false, false, true},
	}

	for _, c := range cases {
		got := biConditional(c.inputP, c.inputQ)
		if got != c.want {
			t.Errorf("biConditional(%t,%t) == %t, want %t", c.inputP, c.inputQ, got, c.want)
		}
	}

}

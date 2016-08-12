package Context

// Context
// Test
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"testing"
)

func TestAddInt(t *testing.T) {
	c := New()
	c.Set("x", 7)
	if c.Get("x") != 7 {
		t.Error("Not added variable of type integer")
	}
}

func TestAddString(t *testing.T) {
	c := New()
	c.Set("x", "s")
	if c.Get("x") != "s" {
		t.Error("Not added variable of type string")
	}
}

func TestChangeType(t *testing.T) {
	c := New()
	c.Set("x", "1")
	if c.Get("x") == 1 {
		t.Error("Function to change the type of the variable")
	}
}

func TestGetNull1(t *testing.T) {
	c := New()
	if c.Get("x") != nil {
		t.Error("Received an undefined variable")
	}
}

func TestGetNull2(t *testing.T) {
	c := New()
	c.Set("x", 7)
	if c.Get("y") != nil {
		t.Error("Received a non-existent variable")
	}
}

func TestGetMidContext(t *testing.T) {
	c := New()
	c.Set("a", 5)
	c.Set("b", 6)
	c.Set("c", 7)
	c.Set("b", 9)
	if c.Get("a") != 5 {
		t.Error("Error in getting the first variable")
	}
	if c.Get("b") != 9 {
		t.Error("Error in getting a modified variable")
	}
}

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

func TestFixContext(t *testing.T) {
	c1 := New()
	c1.Set("a", 5)
	c1.Set("b", 6)
	c1.Set("c", 7)

	c10 := c1.Fix()
	c10.Set("a", 15)
	c10.Set("b", 16)
	c10.Set("c", 17)

	c20 := c1.Fix()
	c20.Set("a", 25)

	if c1.Get("a") != 5 {
		t.Error("Error in getting a modified variable 'a' - 5 != ", c1.Get("a"))
	}
	if c1.Get("b") != 6 {
		t.Error("Error in getting a modified variable 'b' - 6 != ", c1.Get("b"))
	}
	if c1.Get("c") != 7 {
		t.Error("Error in getting a modified variable 'c' - 7 != ", c1.Get("c"))
	}
	if c10.Get("a") != 15 {
		t.Error("Error in getting a modified variable 'a' - 15 != ", c10.Get("a"))
	}
	if c10.Get("b") != 16 {
		t.Error("Error in getting a modified variable 'b' - 16 != ", c10.Get("b"))
	}
	if c10.Get("c") != 17 {
		t.Error("Error in getting a modified variable 'c' - 17 != ", c10.Get("c"))
	}
	if c20.Get("a") != 25 {
		t.Error("Error in getting a modified variable 'a' - 25 != ", c20.Get("a"))
	}
	if c20.Get("b") != 6 {
		t.Error("Error in getting inherited variable 'b' - 6 != ", c20.Get("b"))
	}
}

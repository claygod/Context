package Context

// Context
// Work
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

// Context structure
type Context struct {
	link *Node
}

// New - create a new Context
func New() *Context {
	c := &Context{newNode()}
	return c
}

// Set - add the variable in context.
func (c *Context) Set(key interface{}, value interface{}) {
	c.link = c.link.add(key, value)
}

// Get - get variable.
func (c *Context) Get(key interface{}) interface{} {
	return c.link.search(key)
}

// Node structure
type Node struct {
	parent *Node
	key    interface{}
	value  interface{}
}

// newNode - create a new Node-struct
func newNode() *Node {
	c := &Node{}
	return c
}

// search - find variable in nodes.
func (c *Node) search(key interface{}) interface{} {
	if c.key == key {
		return c.value
	}
	if c.parent != nil {
		return c.parent.search(key)
	}
	return nil
}

// add - add a node to a new variable.
func (c *Node) add(key interface{}, value interface{}) *Node {
	x := newNode()
	x.parent = c
	c.key = key
	c.value = value
	return x
}

package main

type Food struct {
	name string
}

type Dinner struct {
	f []Food
}

func (d *Dinner) CreateIterator() *ConcreteIterator {
	if d == nil {
		return nil
	}

	return &ConcreteIterator{d, 0}
}

func (d *Dinner) Add(f Food) {
	d.f = append(d.f, f)
}

type Iterator interface {
	first() interface{}
	next() interface{}
}

type ConcreteIterator struct {
	dinner *Dinner
	index  int
}

func (c *ConcreteIterator) First() interface{} {
	if c == nil {
		return nil
	}

	if len(c.dinner.f) > 0 {
		c.index = 0
		return c.dinner.f[0]
	}
	return nil
}

func (c *ConcreteIterator) Next() interface{} {
	if c == nil {
		return nil
	}

	if len(c.dinner.f) > c.index+1 {
		c.index++
		return c.dinner.f[c.index]
	}

	return nil
}

package factory

import "fmt"

// Product declares the interface, which is common to all objects that can be produced by the creator and its subclasses.
type Product interface {
	doStuff() string
}

// ConcreteProductA Concrete Products are different implementations of the prod- uct interface.
type ConcreteProductA struct {
}

func (p *ConcreteProductA) doStuff() string {
	fmt.Printf("ConcreteProductA.doStuff()\n")
	return "ConcreteProductA"
}

type ConcreteProductB struct {
}

func (p *ConcreteProductB) doStuff() string {
	fmt.Printf("ConcreteProductB.doStuff()\n")
	return "ConcreteProductB"
}

// The Creator class declares the factory method that returns new product objects. It’s important that the return type of this method matches the product interface.
type Creator struct {
	create ProductCreator
}

type ProductCreator interface {
	createProduct() Product
}

func (c *Creator) someOperation() string {
	fmt.Printf("Creator.someOperation()\n")
	p := c.createProduct()
	return p.doStuff()
}

func (c *Creator) createProduct() Product {
	return c.create.createProduct()
}

type ConcreteCreatorA struct {
}

func (c *ConcreteCreatorA) createProduct() Product {
	return &ConcreteProductA{}
}

type ConcreteCreatorB struct {
}

func (c *ConcreteCreatorB) createProduct() Product {
	return &ConcreteProductB{}
}

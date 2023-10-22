package solid

import "fmt"

// Open-Closed Principle
// open for extension, closed for modification
// Specification Pattern

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

// FilterByColor create a filter that filters by color
func (f *Filter) FilterByColor(
	product []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range product {
		if v.color == color {
			result = append(result, &product[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(
	product []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range product {
		if v.size == size {
			result = append(result, &product[i])
		}
	}
	return result
}

func (f *Filter) FilterBySizeAndColor(
	product []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range product {
		if v.size == size && v.color == color {
			result = append(result, &product[i])
		}
	}
	return result
}

// Specification Pattern
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type BetterFilter struct{}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) &&
		a.second.IsSatisfied(p)
}

func (f *BetterFilter) Filter(
	product []Product,
	spec Specification,
) []*Product {
	result := make([]*Product, 0)
	for i, p := range product {
		if spec.IsSatisfied(&p) {
			result = append(result, &product[i])
		}
	}
	return result
}

func solidOCP() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	filter := Filter{}

	for _, v := range filter.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	// ColorSpecification
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	fmt.Println("Green products (new):")
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	// AndSpecification
	smallSpec := SizeSpecification{small}
	greenAndSmallSpec := AndSpecification{greenSpec, smallSpec}
	fmt.Println("Small and green products:")
	for _, v := range bf.Filter(products, greenAndSmallSpec) {
		fmt.Printf(" - %s is small and green\n", v.name)
	}
}

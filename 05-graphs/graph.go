package graphs

import (
	"fmt"
	"strconv"
)

type Graph struct {
	vertices []*Vertex
}

func NewGraph() *Graph {
	return &Graph{}
}

func (g *Graph) AddVertex(v *Vertex) {
	g.vertices = append(g.vertices, v)
}

func (g *Graph) AddEdge(source, destination *Vertex) error {
	if !g.vertexExists(source) {
		return fmt.Errorf("source vertex does not exist")
	}
	if !g.vertexExists(destination) {
		return fmt.Errorf("destination vertex does not exist")
	}
	if err := source.AddEdge(destination); err != nil {
		return err
	}
	return nil
}

func (g *Graph) vertexExists(v *Vertex) bool {
	for _, vertex := range g.vertices {
		if vertex.key == v.key {
			return true
		}
	}
	return false
}

func (g *Graph) Print() string {
	var result string
	for _, v := range g.vertices {
		result += fmt.Sprintf("[%d]", v.key)
		for _, a := range v.adjacent {
			result += fmt.Sprintf(" -> (%d)", a.key)
		}
		result += "\n"
	}
	return result
}

func (g *Graph) Vertices() int {
	return len(g.vertices)
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func NewVertex(key int) *Vertex {
	return &Vertex{key: key}
}

func (v *Vertex) String() string {
	return strconv.Itoa(v.key)
}

func (v *Vertex) AddEdge(destination *Vertex) error {
	for _, a := range v.adjacent {
		if a.key == destination.key {
			return fmt.Errorf("vertices already connected")
		}
	}
	v.adjacent = append(v.adjacent, destination)
	return nil
}

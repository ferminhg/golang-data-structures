package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Graph(t *testing.T) {
	t.Run("Given an YOLO vertice When we print it Then it's correct", func(t *testing.T) {
		vertice := NewVertex(1)
		result := vertice.String()
		assert.Equal(t, "1", result)

		graph := NewGraph()
		graph.AddVertex(vertice)
		assert.Equal(t, 1, graph.Vertices())
		assert.Equal(t, "[1]\n", graph.Print())
	})

	t.Run("Given a graph with 3 vertices When we print it Then it's correct", func(t *testing.T) {
		graph := NewGraph()
		graph.AddVertex(NewVertex(1))
		graph.AddVertex(NewVertex(2))
		graph.AddVertex(NewVertex(3))
		assert.Equal(t, 3, graph.Vertices())
		assert.Equal(t, "[1]\n[2]\n[3]\n", graph.Print())
	})

	t.Run("Given a graph with 2 vertices When we connect them Then they are connected", func(t *testing.T) {
		graph := NewGraph()
		vertex1 := NewVertex(1)
		vertex2 := NewVertex(2)
		graph.AddVertex(vertex1)
		graph.AddVertex(vertex2)
		graph.AddEdge(vertex1, vertex2)
		assert.Equal(t, 1, len(graph.vertices[0].adjacent))
		assert.Equal(t, 0, len(graph.vertices[1].adjacent))
		assert.Equal(t, "[1] -> (2)\n[2]\n", graph.Print())
	})

	t.Run("Given a graph with 1 vertice When we connect unexisting vertices Then we get an error", func(t *testing.T) {
		graph := NewGraph()
		vertex1 := NewVertex(1)
		graph.AddVertex(vertex1)

		assert.EqualError(t, graph.AddEdge(vertex1, NewVertex(2)), "destination vertex does not exist")
		assert.EqualError(t, graph.AddEdge(NewVertex(2), vertex1), "source vertex does not exist")
	})

	t.Run("Given a graph with 2 vertices When we connect them twice Then we get an error", func(t *testing.T) {
		graph := NewGraph()
		vertex1 := NewVertex(1)
		vertex2 := NewVertex(2)
		graph.AddVertex(vertex1)
		graph.AddVertex(vertex2)
		graph.AddEdge(vertex1, vertex2)
		assert.EqualError(t, graph.AddEdge(vertex1, vertex2), "vertices already connected")
		assert.Equal(t, "[1] -> (2)\n[2]\n", graph.Print())
	})
}

package fun_facts

import "testing"

func Test(t *testing.T) {
	t.Run("deferer", func(t *testing.T) {
		deferer()
	})

	t.Run("copier", func(t *testing.T) {
		copier()
	})

	t.Run("adder", func(t *testing.T) {
		adder()
	})

	t.Run("slicer", func(t *testing.T) {
		slicer()
	})
}

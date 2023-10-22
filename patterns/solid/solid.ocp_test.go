package solid

import "testing"

func Test_solidOCP1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test_solidOCP1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solidOCP()
		})
	}
}

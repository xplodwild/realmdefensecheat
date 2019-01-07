package realmdefense

import "testing"

func TestComputeHash(t *testing.T) {
	tests := []struct {
		name string
		args []byte
		want string
	}{
		{name: "Default test", args: []byte("{}"), want: "152e9aef86eafe7c2a6db2cb91d0c32c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeHash(tt.args); got != tt.want {
				t.Errorf("ComputeHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

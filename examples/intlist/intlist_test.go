package intlist

import "testing"

func TestIntList_Sum(t *testing.T) {
	type fields struct {
		Value int
		Tail  *IntList
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{"test nil", fields{
			Value: 0,
			Tail:  &IntList{},
		}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &IntList{
				Value: tt.fields.Value,
				Tail:  tt.fields.Tail,
			}
			if got := list.Sum(); got != tt.want {
				t.Errorf("IntList.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

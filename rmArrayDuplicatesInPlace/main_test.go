package main

import (
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		newNums []int
	}{
		{name: "happy 1", args: args{nums: []int{1, 1, 2}}, want: 2, newNums: []int{1, 2}},
		{name: "0 length", args: args{nums: []int{}}, want: 0, newNums: []int{}},
		{name: "1 length", args: args{nums: []int{1}}, want: 1, newNums: []int{1}},
		{name: "2 length, 1 dupe", args: args{nums: []int{1, 1}}, want: 1, newNums: []int{1}},
		{name: "3 length, 2 dupes", args: args{nums: []int{1, 1, 1}}, want: 1, newNums: []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)

			if got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}

			for i := 0; i < got; i++ {
				if tt.args.nums[i] != tt.newNums[i] {
					t.Errorf("wanted: %v\n but got: %v\n", tt.want, tt.args.nums)
				}
			}
		})
	}
}

package main

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "happy", args: args{prices: []int{7, 1, 5, 3, 6, 4}}, want: 7},
		{name: "happy2", args: args{prices: []int{1, 2, 3, 4, 5}}, want: 4},
		{name: "happy3", args: args{prices: []int{7, 6, 4, 3, 2, 1}}, want: 0},
		{name: "0 len", args: args{prices: []int{}}, want: 0},
		{name: "happy4", args: args{prices: []int{7, 6, 4, 3, 2, 1, 2}}, want: 1},
		{name: "happy5", args: args{prices: []int{6, 1, 3, 2, 4, 7}}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

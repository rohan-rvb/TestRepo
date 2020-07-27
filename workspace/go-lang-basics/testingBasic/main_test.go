package main

import "testing"

type AddResult struct {
	x int
	y int
	ex int
}

var AddResults = []AddResult{
	{
		x:1,
		y:2,
		ex:3,
	},
	{
		x:5,
		y:8,
		ex:14,
	},
}

func TestAdd(t *testing.T) {
	for _,v := range AddResults {
		if Add(v.x, v.y) != v.ex {
			t.Fatalf("Error occerr")
		}
	}
}

package math

import (
	"testing"
)

type mathTest struct {
	arg1   int
	arg2   int
	result int
}

var dataAdd []mathTest = []mathTest{
	{1, 1, 2},
	{2, 2, 4},
	{5, 5, 10},
	{7, 8, 15},
	// {7, 8, 16},
}

var dataSubstract []mathTest = []mathTest{
	{1, 1, 0},
	{2, 2, 0},
	{5, 5, 0},
	{7, 8, -1},
	// {7, 8, 1},
}

func TestAdd(t *testing.T) {

	for _, val := range dataAdd {
		if Add(val.arg1, val.arg2) != val.result {
			t.Error("error occured for val ", val)
		}
	}
}

func TestSubstract(t *testing.T) {

	for _, val := range dataSubstract {
		if Substract(val.arg1, val.arg2) != val.result {
			t.Error("error occured for val ", val)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}

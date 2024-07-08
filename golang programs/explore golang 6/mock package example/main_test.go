package main

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

// Now for writting tests we will write a struct which will implement all the methods of that interface that we have defined.
type mockRand struct {
	mock.Mock
}

func NewMockRand() *mockRand {
	return &mockRand{}
}

// in order to use the m.On("randomInt", 10).Return(6) . we need to give mockRand a randomInt method. without it can't do that.
func (m *mockRand) randomInt(max int) int {
	returnValues := m.Called(max) // it keeps record that it was called with the value that passed in for max // it also returns the Arguments object which contains the return values that we have specified in the test to returned if the function gets the given value of max. NOTE : here max is a variable.
	return returnValues.Int(0)    // here we are telling that return the zeroth value only.
}

// In addition to arguments.Int there is Bool, String, Error methods available for the same.

// if we have some other type that we have method Arguments.Get, which returns an interface{}

// Now the steps for test are

// create an instance of mock
// specify what results you want back when mocks methods are called. using On and Return methods
// run the code that is being tested, the standard way you would do in Go Test.
// optionally use the methods like Mock.AssertCalled to check that a given method indeed has been called during the test.

func TestDivByRand(t *testing.T) {
	m := NewMockRand()

	// when randomInt get called with value 10 then return 20 value from it.
	m.On("randomInt", 10).Return(5)
	// m.On("randomInt", 20).Return(40)

	// now do assertions as we would do in standard go testing

	// NOTE: HERE 10 act as a seed value

	// if i don't have any output for seed value then i will get the unexpected method call error
	// hence we should pass the possible arguments and possible return values to the functions so as to mimic it.
	quotient := divByRand(40, 10, m)

	if quotient != 8 {
		t.Errorf("expected quotient to be 5, got %d", quotient)
	}

	// check that randomInt was called with the number 10
	// if not then the test fails
	m.AssertCalled(t, "randomInt", 10)
}

//

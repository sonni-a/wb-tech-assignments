package main

import (
	"reflect"
	"testing"
)

func TestGroupTemps(t *testing.T) {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	expected := map[int][]float64{
		-20: {-25.4, -27.0, -21.0},
		10:  {13.0, 19.0, 15.5},
		20:  {24.5},
		30:  {32.5},
	}

	result := GroupTemps(temps)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got: %v\n want: %v", result, expected)
	}
}

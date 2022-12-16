package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testValues = []struct {
	input  string
	expect uint64
}{
	{"0.01", 1},
	{"99.99", 9999},
	{"неЧисло", 0},
	{"100", 0},
	{"-2", 0},
	{"5,29", 529},
	{"0,01", 1},
	{"99,99", 9999},
	{"2,,0", 0},
	{"1000,29", 0},
	{"-12,25", 0},
	{"53.999", 0},
	{"26.1", 2610},
	{"44,44", 4444},
	{"0,00", 0},
	{"1,00", 100},
	{"1,01", 101},
	{"1.5", 150},
	{"1,5", 150},
	{" ", 0},
	{"1 ,02", 102},
	{" 0,5", 50},
	{"", 0},
	{"88,3 ", 8830},
	{" 0,29 ", 29},
	{" 1 , 2 9   \n", 129},
	{"\n\t \n1\n\t2.\t33   ", 1233},
	{".12", 12},
	{"55.", 5500},
	{"28.2.3", 0},
	{"53.2.", 0},
	{"NaN", 0},
}

func TestCheckComission(t *testing.T) {
	for _, testValue := range testValues {
		res := check_commission(testValue.input)
		assert.Equal(t, res, testValue.expect)
	}
}

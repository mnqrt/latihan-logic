package lat3

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNo1(t *testing.T) {
	expected := [][]int{{1, 0, 0}, {5, 3, 0}, {7, 9, 11}}
	result := No1(3)
	assert.Equal(t, expected, result)
}

func TestNo2(t *testing.T) {
	expected := [][]int{{1, 3, 5}, {0, 9, 7}, {0, 0, 11}}
	result := No2(3)
	assert.Equal(t, expected, result)
}

func TestNo3(t *testing.T) {
	expected := [][]int{{1, 4, 7}, {13, 10, 0}, {15, 0, 0}}
	result := No3(3)
	assert.Equal(t, expected, result)
}

func TestNo4(t *testing.T) {
	expected := [][]int{{0, 0, 1}, {0, 3, 5}, {11, 9, 7}}
	result := No4(3)
	assert.Equal(t, expected, result)
}

func TestNo5(t *testing.T) {
	expected := [][]int{{1, 0, 1}, {5, 3, 5}, {1, 0, 1}}
	result := No5(3)
	assert.Equal(t, expected, result)
}

func TestNo6(t *testing.T) {
	expected := [][]int{{1, 3, 5}, {0, 7, 0}, {1, 3, 5}}
	result := No6(3)
	assert.Equal(t, expected, result)
}

func TestNo7(t *testing.T) {
	expected := [][]int{{0, 1, 0}, {3, 5, 7}, {0, 1, 0}}
	result := No7(3)
	assert.Equal(t, expected, result)
}

func TestNo8(t *testing.T) {
	expected := [][]int{{0, 3, 0}, {1, 5, 1}, {0, 7, 0}}
	result := No8(3)
	assert.Equal(t, expected, result)
}

func TestNo9(t *testing.T) {
	expected := [][]int{{0, 1, 0}, {1, 3, 1}, {0, 1, 0}}
	result := No9(3)
	assert.Equal(t, expected, result)
}

func TestNo10(t *testing.T) {
	expected := [][]int{{4, 3, 4}, {3, 1, 3}, {4, 3, 4}}
	result := No10(3)
	assert.Equal(t, expected, result)
}

func TestNo11(t *testing.T) {
	expected := [][]int{{1, 0, 1}, {0, 3, 5}, {1, 0, 1}}
	result := No11(3)
	assert.Equal(t, expected, result)
}

func TestNo12(t *testing.T) {
	expected := [][]int{{1, 0, 1}, {5, 3, 0}, {1, 0, 1}}
	result := No12(3)
	assert.Equal(t, expected, result)
}

func TestNo12B(t *testing.T) {
	expected := [][]int{{0, 1, 0, 1}, {3, 5, 3, 0}, {0, 1, 0, 1}}
	result := No12B(3)
	assert.Equal(t, expected, result)
}

func TestNo14(t *testing.T) {
	expected := [][]int{{1, 8, 13}, {1, 5, 21}, {2, 3, 34}}
	result := No14(3)
	assert.Equal(t, expected, result)
}

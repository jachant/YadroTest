package tests

import (
	"fmt"
	"testing"

	"github.com/jachant/YadroTest/task"
)

func TestTask1_Base(t *testing.T) {
	input := [][]uint64{
		{10, 20, 30},
		{1, 1, 1},
		{0, 0, 1},
	}
	expected := "no"
	result := task.Task1(input)
	fmt.Println("input=", input)
	fmt.Println("expected=", expected)
	fmt.Println("result=", result)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

}
func TestTask1_Success(t *testing.T) {
	input := [][]uint64{
		{1, 2},
		{2, 1},
	}
	expected := "yes"
	result := task.Task1(input)
	fmt.Println("input=", input)
	fmt.Println("expected=", expected)
	fmt.Println("result=", result)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestTask1_Failure(t *testing.T) {
	input := [][]uint64{
		{1, 2},
		{3, 4},
	}
	expected := "no"
	result := task.Task1(input)
	fmt.Println("input=", input)
	fmt.Println("expected=", expected)
	fmt.Println("result=", result)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestTask1_SingleContainer(t *testing.T) {
	input := [][]uint64{{5}}
	expected := "yes"
	result := task.Task1(input)
	fmt.Println("input=", input)
	fmt.Println("expected=", expected)
	fmt.Println("result=", result)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestTask1_SortedSums(t *testing.T) {
	input := [][]uint64{
		{5, 0},
		{0, 3},
	}
	expected := "yes"
	result := task.Task1(input)
	fmt.Println("input=", input)
	fmt.Println("expected=", expected)
	fmt.Println("result=", result)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestTask1_AllZeros(t *testing.T) {
	input := [][]uint64{
		{0, 0},
		{0, 0},
	}
	expected := "yes"
	result := task.Task1(input)
	fmt.Println("input=", input)
	fmt.Println("expected=", expected)
	fmt.Println("result=", result)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestTask1_DuplicateSums(t *testing.T) {
	input := [][]uint64{
		{3, 3},
		{3, 3},
	}
	expected := "yes"
	result := task.Task1(input)
	fmt.Println("input=", input)
	fmt.Println("expected=", expected)
	fmt.Println("result=", result)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestTask1_NonSquareMatrix(t *testing.T) {
	input := [][]uint64{
		{1, 2, 3},
		{4, 5, 6},
	}
	expected := "no"
	result := task.Task1(input)
	fmt.Println("input=", input)
	fmt.Println("expected=", expected)
	fmt.Println("result=", result)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

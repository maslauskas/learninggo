package main

import (
    "fmt"
    "reflect"
    "testing"
)

func TestAdd(t *testing.T) {
    t.Run("it adds 2 and 2", func(t *testing.T) {
        actual := Add(2, 2)
        expected := 4

        assertNumberEquals(t, actual, expected)
    })

    t.Run("it adds numbers from collection of fixed size", func(t *testing.T) {
        numbers := []int{1,2,3,4,5}
        actual := Sum(numbers)
        expected := 15

        assertNumberEquals(t, actual, expected)
    })

    t.Run("it adds numbers from collection of any size", func(t *testing.T) {
        numbers := []int{1,2,3,4,5,6}
        actual := Sum(numbers)
        expected := 21

        assertNumberEquals(t, actual, expected)
    })
}

func TestSumAll(t *testing.T) {
    actual := SumAll([]int{1,2}, []int{0,9})
    expected := []int{3,9}

    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("expected result does not match, %v", actual)
    }
}

func TestSumAllTails(t *testing.T) {
    t.Run("with empty input", func(t *testing.T) {
        actual := SumAllTails([]int{})
        expected := []int{0}

        if !reflect.DeepEqual(actual, expected) {
            t.Errorf("expected %v, got %v", expected, actual)
        }
    })

    t.Run("with example input", func(t *testing.T) {
        actual := SumAllTails([]int{1,2}, []int{0,9})
        expected := []int{2, 9}

        if !reflect.DeepEqual(actual, expected) {
            t.Errorf("expected %v, got %v", expected, actual)
        }
    })
}

func ExampleAdd() {
    result := Add(3,6)
    fmt.Println(result)
    // Output: 9
}

func assertNumberEquals(t *testing.T, actual int, expected int) {
    if actual != expected {
        t.Errorf("expected %d, got %d", expected, actual)
    }
}
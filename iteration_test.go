package main

import (
    "strings"
    "testing"
)

func TestIteration(t *testing.T) {
    t.Run("will print empty string when repeat is 0", func(t *testing.T) {
        actual := Repeat("x", 0)
        expected := ""

        assertStringEquals(t, actual, expected)
    })

    t.Run("will print x when repeat is 1", func(t *testing.T) {
        actual := Repeat("x", 1)
        expected := "x"

        assertStringEquals(t, actual, expected)
    })

    t.Run("will print xxxx when repeat is 4", func(t *testing.T) {
        actual := Repeat("x", 4)
        expected := "xxxx"

        assertStringEquals(t, actual, expected)
    })
}

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("x", i)
    }
}

func Repeat(input string, times int) string {
    return strings.Repeat(input, times)
}
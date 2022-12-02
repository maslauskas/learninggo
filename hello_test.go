package main

import "testing"

func TestHello(t *testing.T) {
    t.Run("will print hello name when name is provided", func (t *testing.T) {
        actual := hello("Tadas")
        expected := "Hello, Tadas"

        assertStringEquals(t, actual, expected)
    })

    t.Run("will print hello world when empty string is provided", func (t *testing.T) {
        actual := hello("")
        expected := "Hello, world"

        assertStringEquals(t, actual, expected)
    })
}

func assertStringEquals(t *testing.T, actual string, expected string) {
    if actual != expected {
        t.Errorf("expected %q, got %q", expected, actual)
    }
}

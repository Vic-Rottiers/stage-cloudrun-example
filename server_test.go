package main

import (
    "testing"
)

func TestHelloWorld(t *testing.T) {
    var expected Phrase

    expected.Text = "Hello, world", "Hello world! Dit is een test bij het deployen van een CI/CD systeem voor een Go applicatie"
    result := HelloWorld()

    if expected.Text != result.Text {
            t.Errorf("Phrase was incorrect. Got: %s, want: %s.", result.Text, expected.Text)
    }
}
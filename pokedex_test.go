package main

import (
  "testing"
)

func TestCleanInput(t *testing.T) {
  // test cases
  cases := []struct{
    input string
    expected []string
  }{
    {
      input: " hello world ",
      expected: []string{"hello", "world"},
    },
    {
      input: "         hello",
      expected: []string{"hello"},
    },
    {
      input: "hello hello  hello  ",
      expected: []string{"hello", "hello", "hello"},
    },
  }
  // run cases
  for _, c := range cases {
    actual := cleanInput(c.input)
    if len(actual) != len(c.expected) {
      t.Errorf("Length of actual: %d and Length of Expected: %d does not match!", len(actual), len(c.expected))
    } else {
      for i := range actual {
        word := actual[i]
        expectedWord := c.expected[i]
        if word != expectedWord {
          t.Errorf("Word #%d: %s does not match %s", i, word, expectedWord)
        }
      }
    }
  }
}

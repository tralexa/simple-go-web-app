package main

import "testing"

func TestThis1(t *testing.T) {
  t.Skip("skipping test in short mode.")
}

func TestThis2(t *testing.T) {
  //t.Skip("skipping test in short mode.")
}

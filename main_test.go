package main

import (
          "testing"
)

func TestMakeGreetin(t *testing.T) {
          want := "Hello, Taro"
          got := makeGreetin("Taro")
          if got != want {
                    t.Errorf("got = %s; want %s", got, want)
          }
}

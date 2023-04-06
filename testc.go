package testc

import "fmt"

func Hello(s string) string {
    res := fmt.Sprintf("hello,%s", s)
    return res
}

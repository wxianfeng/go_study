package main

import (
	"fmt"
)

func max(a int, b int) (int) {
	if a > b {
        return a
    }
        
     return b   
}

func Add(a int) (int) {
    a = a + 1

    return a
}

func AddByPointer(a *int) (int) {
    *a = *a + 1

    return *a
}

func main() {
    max_value := max(2, 3)
    fmt.Printf("max value: %d\n", max_value)

    a := 100
    incr_a := Add(a)
    fmt.Printf("Add by value, a +1 is: %d\n", incr_a) // 101
    fmt.Printf("Add by value, a is: %d\n", a) // 100, 值传递, 原始的没有被修改

    a2 := 100
    incr_a2 := AddByPointer(&a2)
    fmt.Printf("AddByPointer, a2 +1 is: %d\n", incr_a2) // 101
    fmt.Printf("AddByPointer a2 is: %d\n", a2) // 101, 指针传递, 原始值被修改
}

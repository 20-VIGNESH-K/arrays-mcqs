// package main

// import (
//     "fmt"
// )

// func hello() []string {
//     return nil
// }

// func main() {
//     h := hello
//     if h == nil {
//         fmt.Println("nil")
//     } else {
//         fmt.Println("not nil")
//     }
// }

// Options
// nil
// not nil
// compilation error








// package main

// import (
//     "fmt"
//     "strconv"
// )

// func main() {
//     i := 2
//     s := "1000"
//     if len(s) > 1 {
//         i, _ := strconv.Atoi(s)
//         i = i + 5
//     }
//     fmt.Println(i)
// }

// Options
// 2
// 1005
// compilation error








// package main

// import (
//     "fmt"
// )

// func hello(num ...int) {
//     num[0] = 18
// }

// func main() {
//     i := []int{5, 6, 7}
//     hello(i...)
//     fmt.Println(i[0])
// }

// Options
// 18
// 5
// Compilation error










// package main

// import (
//     "fmt"
// )

// func main() {
//     a := [2]int{5, 6}
//     b := [2]int{5, 6}
//     if a == b {
//         fmt.Println("equal")
//     } else {
//         fmt.Println("not equal")
//     }
// }

// Options
// compilation error
// equal
// not equal










// package main

// import "fmt"

// type rect struct {
//     len, wid int
// }

// func (r rect) area() {
//     fmt.Println(r.len * r.wid)
// }

// func main() {
//     r := &rect{len: 5, wid: 6}
//     r.area()
// }

// Options
// compilation error
// 30











// package main

// import (
//     "fmt"
// )

// func main() {
//     a := [5]int{1, 2, 3, 4, 5}
//     t := a[3:4:4]
//     fmt.Println(t[0])
// }

// Options
// 3
// 4
// compilation error












// package main

// import (
//     "fmt"
// )

// type person struct {
//     name string
// }

// func main() {
// var m map[string]int
// m := make(map[string]float64)
// m["mike"]=1
// p := "mike"
// fmt.Println(m["a"])
// }

// Options
// compilation error
// 0
// 1








// package main

// import (
//     "fmt"
// )

// func main() {
//     var i interface{}

//     if i == nil {
//         fmt.Println("nil")
//         return
//     }
//     fmt.Println("not nil")
// }

// Options
// nil
// not nil
// compilation error
















// package main

// import (
//     "fmt"
// )

// func main() {
//     s := make(map[string]int)
//     delete(s, "h")
//     fmt.Println(s["h"])
// }

// Options
// runtime panic
// 0
// compilation error




















// package main

// import "fmt"

// func main() {
// 	var x = map[string]int{"h": 1, "y": 2}
// 	delete(x, "h")
// 	fmt.Println(x)
// }

// A) {“a”: 1}
// B) {“b”: 2}
// C) {}
// D) Error

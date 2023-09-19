// Code:
// package main

// import "fmt"

// func main() {
//     slice := []int{1, 2, 3}
//     fmt.Println(len(slice))
// }
// MCQ: What will be the output of this code? a) 1 b) 2 c) 3 d) 0










// Code:
// package main

// import "fmt"

// func main() {
//     slice := make([]int, 5)
//     fmt.Println(cap(slice))
// }
// MCQ: What will be the output of this code? a) 0 b) 5 c) 1 d) Error












// Code:
// package main

// import "fmt"

// func main() {
//     slice := []string{"apple", "banana", "cherry"}
//     fmt.Println(slice[1])
// }
// MCQ: What will be the output of this code? a) apple b) banana c) cherry d) Error













// Code:
// package main

// import "fmt"

// func main() {
//     slice := []bool{true, false, true}
//     fmt.Println(slice[3])
// }
// MCQ: What will be the output of this code? a) true b) false c) Error d) 0














// Code:
// package main

// import "fmt"

// func main() {
//     slice := []int{1, 2, 3, 4, 5}
//     fmt.Println(slice[1:3])
// }
// MCQ: What will be the output of this code? a) [2 3] b) [1 2] c) [3 4] d) Error


















// Code:
// package main

// import "fmt"

// func main() {
//     slice := []int{1, 2, 3}
//     slice = append(slice, 4)
//     fmt.Println(slice)
// }
// MCQ: What will be the output of this code? a) [1 2 3 4] b) [4 3 2 1] c) [1 2 3] d) Error


















// Code:
// package main

// import "fmt"

// func main() {
//     slice := []int{1, 2, 3}
//     newSlice := slice[1:3]
//     fmt.Println(newSlice[1])
// }
// MCQ: What will be the output of this code? a) 1 b) 2 c) 3 d) Error

















// Code:
// package main

// import "fmt"

// func main() {
//     var slice []int
//     fmt.Println(len(slice))
// }
// MCQ: What will be the output of this code? a) 0 b) 1 c) Error d) nil
















// Code:
// package main

// import "fmt"

// func main() {
//     slice1 := []int{1, 2, 3}
//     slice2 := slice1
//     slice2[0] = 4
//     fmt.Println(slice1[0])
// }
// MCQ: What will be the output of this code? a) 1 b) 2 c) 3 d) 4























// Code:
// package main

// import "fmt"

// func main() {
//     slice := []int{1, 2, 3, 4, 5}
//     slice = slice[:2]
//     fmt.Println(len(slice))
// }
// MCQ: What will be the output of this code? a) 1 b) 2 c) 3 d) 5
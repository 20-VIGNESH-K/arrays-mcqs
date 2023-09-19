// // package main

// // import "fmt"











// // func main() {
// // 	var numbers [5]int
// // 	numbers = [5]int{1, 2, 3, 4, 5}
// // 	fmt.Printf("%d\n", numbers[3])
// // }

// // What is the output of this program?
// // a) 1
// // b) 2
// // c) 3
// // d) 4





















// // func main() {
// //     numbers := [5]int{1, 2, 3, 4, 5}
// //     sum := 0
// //     for i := 0; i < len(numbers); i++ {
// //         sum += numbers[i]
// //     }
// //     fmt.Printf("Sum: %d\n", sum)
// // }
// // What is the output of this program?
// // a) 5
// // b) 10
// // c) 15
// // d) 25



















// // func main() {
// //     source := [5]int{1, 2, 3, 4, 5}
// //     var target [5]int
// //     for i := 0; i < len(source); i++ {
// //         target[i] = source[i]
// //     }
// //     fmt.Printf("%d\n", target[2])
// // }
// // What is the output of this program?
// // a) 1
// // b) 2
// // c) 3
// // d) 4















// package main

// import "fmt"

// func main() {
//     arr := [3]int{1, 2, 3}
//     fmt.Println(arr[1])
// }
// // MCQ: What will be the output of this code? a) 1 b) 2 c) 3 d) Error
















// // Code:
// package main

// import "fmt"

// func main() {
//     var arr [5]int
//     fmt.Println(len(arr))
// }
// // MCQ: What will be the output of this code? a) 0 b) 5 c) 1 d) Error
















// // Code:
// package main

// import "fmt"

// func main() {
//     arr := [3]string{"apple", "banana", "cherry"}
//     fmt.Println(arr[2])
// }
// // MCQ: What will be the output of this code? a) apple b) banana c) cherry d) Error



















// // Code:
// package main

// import "fmt"

// func main() {
//     arr := [4]bool{true, false, true, true}
//     fmt.Println(arr[3])
// }
// // MCQ: What will be the output of this code? a) true b) false c) Error d) 0



















// // Code:
// package main

// import "fmt"

// func main() {
//     arr := [2][3]int{{1, 2, 3}, {4, 5, 6}}
//     fmt.Println(arr[1][1])
// }
// // MCQ: What will be the output of this code? a) 1 b) 2 c) 5 d) 6

















// // Code:
// package main

// import "fmt"

// func main() {
//     arr := [3]int{1, 2, 3}
//     arr[1] = 5
//     fmt.Println(arr)
// }
// // MCQ: What will be the output of this code? a) [1 5 3] b) [1 2 3] c) [5 5 5] d) Error


















// // Code:
// package main

// import "fmt"

// func main() {
//     arr := [4]int{10, 20, 30, 40}
//     fmt.Println(arr[4])
// }
// // MCQ: What will be the output of this code? a) 10 b) 20 c) 30 d) Error
















// // Code:
// package main

// import "fmt"

// func main() {
//     arr := [3]int{1, 2, 3}
//     arr = [3]int{4, 5, 6}
//     fmt.Println(arr)
// }
// // MCQ: What will be the output of this code? a) [1 2 3] b) [4 5 6] c) [1 4 5] d) Error



















// // Code:
// package main

// import "fmt"

// func main() {
//     arr := [4]int{1, 2, 3, 4}
//     slice := arr[1:3]
//     fmt.Println(len(slice))
// }
// // MCQ: What will be the output of this code? a) 1 b) 2 c) 3 d) Error

















// // Code:
// package main

// import "fmt"

// func main() {
//     arr1 := [3]int{1, 2, 3}
//     arr2 := arr1
//     arr2[0] = 4
//     fmt.Println(arr1[0])
// }
// // MCQ: What will be the output of this code? a) 1 b) 2 c) 3 d) 4
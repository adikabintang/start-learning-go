# What's the difference between array and slice?
read this:
- https://blog.golang.org/go-slices-usage-and-internals
- https://tour.golang.org/moretypes/7
- https://tour.golang.org/moretypes/11

Differences:
- Array has fixed length, slice has dynamic length
- Passing array to functions is "pass by value", slice is "pass by reference"

How to use it:

Declaring:
```go
// array
arr := [3]int{6, 7, 8} // see, fixed value of 3 in [3]

// slice
sl := []int{0, 1, 2} // no value in [].

// declaring empty slice
emptySl := []int{0, 0}
alsoEmptySl := make([]int, 2)

// we can also create a slice from an existing array
sliceFromArray := arr[1:2]
```
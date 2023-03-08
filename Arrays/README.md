In Go language, arrays are a fixed-size sequence of elements of the same type. Here are some of the methods that can be used with arrays in Go:

1.Accessing elements:

arr[index] - retrieves the element at the given index
len(arr) - returns the length of the array
2.Modifying elements:

arr[index] = value - updates the element at the given index with the new value
3.Copying arrays:

copy(dest, src) - copies the elements from the source array to the destination array
4.Slicing arrays:

arr[start:end] - creates a new array containing elements from the start index up to, but not including, the end index
arr[start:] - creates a new array containing elements from the start index to the end of the array
arr[:end] - creates a new array containing elements from the beginning of the array up to, but not including, the end index

Note that arrays in Go are not resizable, so methods like append and remove are not available for arrays. If you need a resizable sequence of elements, consider using a slice instead.
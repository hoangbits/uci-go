### Instructions
The goal of this activity is to explore the use of threads by creating a program for sorting integers that uses four goroutines to create four sub-arrays and then merge the arrays into a single array.

### Detail Instruction
Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

### Review criterialess 
Students will receive five points if the program sorts the integers and five additional points if there are four goroutines that each prints out a set of array elements that it is storing.
package main

// Here is a quick lookup table for various sorting algorithms.
// In practice, when in doubt go with quicksort, or mergesort.
// For distributed map-reduce-like sifting, prefer a mergesort variant.
//
// Algorithm     | Best Case | Worst Case | Avg. Case | Space Cmplx | Stable? |
// --------------+-----------+------------+-----------+-------------+---------+
// MergeSort     | O(nlg(n)) | O(nlg(n))  | O(nlg(n)) | O(n)        | Yes     |
// --------------+-----------+------------+-----------+-------------+---------+
// InsertionSort | O(n)      | O(n^2)     | O(n^2)    | O(1)        | Yes     |
// --------------+-----------+------------+-----------+-------------+---------+
// QuickSort     | O(nlg(n)) | O(n^2)     | O(nlg(n)) | b:logn av:n | No*     |
// --------------+-----------+------------+-----------+-------------+---------+
// HeapSort      | O(nlg(n)) | O(nlg(n))  | O(nlg(n)) | O(1)        | No      |
// --------------+-----------+------------+-----------+-------------+---------+
// CountingSort  | O(k+n)    | O(k+n)     | O(k+n)    | O(k+n)      | Yes     |
// --------------+-----------+------------+-----------+-------------+---------+
//
// * Usually not: There are stable variants too.
//
// Notes:
// - If you need an in-place sorting, then InsertionSort, or HeapSort can be used.
// - The worst-case is so rare, that QuickSort generally performs the best; you can
//     randomize the pivot to almost eliminate the chance of hitting a worst-case
//     scenario.
// - MergeSort is probably the easiest to implement.
// - InsertionSort is generally *not* efficient.
// - You will only see “BubbleSort” in interviews; it is so bad performance-wise,
//     that is is never used, ever.
// - Unless in an interview, don’t reinvent sorting; chances are that your
//     programming language of choice has either a builtin sorting package,
//     or there are reusable libraries that you can link and use.
//     Don’t reinvent the wheel; use battle-tested libraries.
// - There are hybrid “clever” algorithms that use (for example) a stable QuickSort
//     for small input sizes; and switches to an unstable QuickSort or MergeSort
//     for larger sizes of input. (v8’s (Google Chrome) JavaScript Array.prototype.sort
//     method works this way for example).
func main() {
}

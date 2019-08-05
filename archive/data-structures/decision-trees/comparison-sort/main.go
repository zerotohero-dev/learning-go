package main

// In “comparison sort”; the sorted order is determined by a sequence
// of comparisons between pairs of elements.
//
// * InsertionSort
// * SelectionSort
// * BubbleSort
// * QuickSort
// * MergeSort
// * And HeapSort are comparison sorts.
//
// A decision tree represents every sequence an algorithm might make on an
// input of size N.
//
// Using decision trees you can show that **every** comparison sort has a
// lower bound of Omega(n*log(n)). — In other words, unless there is a special
// constraint, none of the above algorihtms can run faster than n*log(n).
//
// Here is a sample decision tree:
//
//            {abc, bca, acb, cab, bac, cba}
//                        |
//                       a<b?
//                      /    \
//         {bac,bca,cba}      {abc, acb, cab}
//               |                   |
//              a<c?                b<c?
//             /   \               /   \
//   {bca, cba}    {bac} {acb, cab}    {abc}
//        |                  |
//       b<c?               a<c?
//      /   \              /   \
// {cba}    {bca}     {cab}    {acb}
//
// Each leaf of the tree is a unique permutation of the input sequence.
// So for a sequence of n items, there are n! leaves of the tree. [1]
//
// The height of the tree, is the height of its root node; and it is the
// number of edges from the root node to the deepest leaf. [2]
//
// A binary tree of height h, can have at most (2^(h+1))-1 nodes. [3]
//
// => (2^(h+1)) ≥ n!
//
// => h+1 ≥ log(n!) = log(1*2*3…*n) = log(1) + log(2) + … + log(n)
//                                  = log(1) + … + log(n/2) + … + log(n)
//                                  //               | n/2 items > log(n/2)
//                                  // n/2 items < log(n/2)
//
// => h+1 ≥ (n/2)log(n/2)
//
// => h ∈ Omega(n*log(n)) [4]
//
// ∴ The depth of a decision tree is lower bound by Omega(n*log(n)) [5]
//
// When we look at the decision tree of a comparison sort on an input of size n:
//
// * Each path from the root the the leafs is one possible sequence of comparisons.
// * Length of the path is the nubmer of comparisons for that instance.
// * ∴ Height of the tree gives the worst-case sequence of comparisons.
// * Height of the tree is Omega(n*log(n)), from [5].
//
// ∴ Every comparison sort requires Omega(n*log(n)) comparisons, in the worst case.
//
// In other words, for comparison sort, in the worst case, you’ll have to make
// at least n*log(n) comparisons.
//
// When you consider the relationship between big Omega and Big Oh; we can
// state that the “worst case” performance of “any” comparison sort cannot be
// better than O(n*log(n))
func main() {
}

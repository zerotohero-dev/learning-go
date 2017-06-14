## Ad-Hoc Notes

* All files in a package live in a single directory.
* Keep package names short and meaningful; don’t use underscores, don’t use compound words.
* The name of a package is part of its type and function names.
* All your Go code is kept in a “workspace”.
* `echo $GOPATH` is your workspace.
* `go get github.com/golang/example/hello`
* `go install github.com/golang/example/hello`
* Don’t deviate from the directory structure; `src`, `bin` and `pkg` folders in `$GOPATH`.
* GOPATH=$HOME is convenient because $HOME/bin is already in your $PATH.

## Queue

* kd-tree
* multi-way merge sort
* binary search
* counting occurrences
* square root
* shell sort
* nut and bolt sorting
* Find median of an array
* In place arrangement of array where negatives come before positives
* Merge sort
* Find second largest key among n keys
* Linear programming
* Shortest path
* Longest common subsequence
* Topological sorting
* Convex hull
* NP-Completeness
* Travelling Salesman

for j = 2 to A.length
    key = A[j]
    i = j-1
    while i > 0 and A[i] > key
        A[i+1] = A[i]
        i--
    A[i+1] = key
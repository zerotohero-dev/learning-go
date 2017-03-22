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

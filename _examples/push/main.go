package main

import (
	"github.com/avdkp/go-git/gitpkg"
	"os"

	. "github.com/avdkp/go-git/_examples"
)

// Example of how to open a repository in a specific path, and push to
// its default remote (origin).
func main() {
	CheckArgs("<repository-path>")
	path := os.Args[1]

	r, err := gitpkg.PlainOpen(path)
	CheckIfError(err)

	Info("git push")
	// push using default options
	err = r.Push(&gitpkg.PushOptions{})
	CheckIfError(err)
}

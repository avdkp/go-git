package main

import (
	"fmt"
	"github.com/avdkp/go-git/gitpkg"
	"os"

	. "github.com/avdkp/go-git/_examples"
	"github.com/avdkp/go-git/plumbing/object"
)

// Open an existing repository in a specific folder.
func main() {
	CheckArgs("<path>")
	path := os.Args[1]

	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := gitpkg.PlainOpen(path)
	CheckIfError(err)

	// Length of the HEAD history
	Info("git rev-list HEAD --count")

	// ... retrieving the HEAD reference
	ref, err := r.Head()
	CheckIfError(err)

	// ... retrieves the commit history
	cIter, err := r.Log(&gitpkg.LogOptions{From: ref.Hash()})
	CheckIfError(err)

	// ... just iterates over the commits
	var cCount int
	err = cIter.ForEach(func(c *object.Commit) error {
		cCount++

		return nil
	})
	CheckIfError(err)

	fmt.Println(cCount)
}

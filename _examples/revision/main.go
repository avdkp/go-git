package main

import (
	"fmt"
	"github.com/avdkp/go-git/gitpkg"
	"os"

	. "github.com/avdkp/go-git/_examples"
	"github.com/avdkp/go-git/plumbing"
)

// Example how to resolve a revision into its commit counterpart
func main() {
	CheckArgs("<path>", "<revision>")

	path := os.Args[1]
	revision := os.Args[2]

	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := gitpkg.PlainOpen(path)
	CheckIfError(err)

	// Resolve revision into a sha1 commit, only some revisions are resolved
	// look at the doc to get more details
	Info("git rev-parse %s", revision)

	h, err := r.ResolveRevision(plumbing.Revision(revision))

	CheckIfError(err)

	fmt.Println(h.String())
}

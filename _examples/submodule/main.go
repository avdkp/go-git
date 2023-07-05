package main

import (
	"github.com/avdkp/go-git/gitpkg"
	"os"

	. "github.com/avdkp/go-git/_examples"
)

// Basic example of how to clone a repository including a submodule and
// updating submodule ref
func main() {
	CheckArgs("<url>", "<directory>", "<submodule>")
	url := os.Args[1]
	directory := os.Args[2]
	submodule := os.Args[3]

	// Clone the given repository to the given directory
	Info("git clone %s %s --recursive", url, directory)

	r, err := gitpkg.PlainClone(directory, false, &gitpkg.CloneOptions{
		URL:               url,
		RecurseSubmodules: gitpkg.DefaultSubmoduleRecursionDepth,
	})

	CheckIfError(err)

	w, err := r.Worktree()
	if err != nil {
		CheckIfError(err)
	}

	sub, err := w.Submodule(submodule)
	if err != nil {
		CheckIfError(err)
	}

	sr, err := sub.Repository()
	if err != nil {
		CheckIfError(err)
	}

	sw, err := sr.Worktree()
	if err != nil {
		CheckIfError(err)
	}

	Info("git submodule update --remote")
	err = sw.Pull(&gitpkg.PullOptions{
		RemoteName: "origin",
	})
	if err != nil {
		CheckIfError(err)
	}
}

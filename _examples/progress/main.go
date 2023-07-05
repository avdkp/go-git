package main

import (
	"github.com/avdkp/go-git/gitpkg"
	"os"

	. "github.com/avdkp/go-git/_examples"
)

// Example of how to show the progress when you do a basic clone operation.
func main() {
	CheckArgs("<url>", "<directory>")
	url := os.Args[1]
	directory := os.Args[2]

	// Clone the given repository to the given directory
	Info("git clone %s %s", url, directory)

	_, err := gitpkg.PlainClone(directory, false, &gitpkg.CloneOptions{
		URL:   url,
		Depth: 1,

		// as git does, when you make a clone, pull or some other operations the
		// server sends information via the sideband, this information can being
		// collected providing a io.Writer to the CloneOptions options
		Progress: os.Stdout,
	})

	CheckIfError(err)
}

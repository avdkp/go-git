package main

import (
	"github.com/avdkp/go-git/gitpkg"
	"log"
	"os"

	"github.com/avdkp/go-git/config"
	"github.com/avdkp/go-git/storage/memory"

	. "github.com/avdkp/go-git/_examples"
)

// Retrieve remote tags without cloning repository
func main() {
	CheckArgs("<url>")
	url := os.Args[1]

	Info("git ls-remote --tags %s", url)

	// Create the remote with repository URL
	rem := gitpkg.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{url},
	})

	log.Print("Fetching tags...")

	// We can then use every Remote functions to retrieve wanted information
	refs, err := rem.List(&gitpkg.ListOptions{
		// Returns all references, including peeled references.
		PeelingOption: gitpkg.AppendPeeled,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Filters the references list and only keeps tags
	var tags []string
	for _, ref := range refs {
		if ref.Name().IsTag() {
			tags = append(tags, ref.Name().Short())
		}
	}

	if len(tags) == 0 {
		log.Println("No tags!")
		return
	}

	log.Printf("Tags found: %v", tags)
}

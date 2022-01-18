package main

import (
	"log"
	"os"
	"strings"

	"github.com/Scribbd/publish-md-to-gdocs-action/api"
)

const (
	TEMPLATE_ENV_KEY    = "TEMPLATE_URL"
	DESTINATION_ENV_KEY = "DESTINATION_URL"
	SOURCE_ENV_KEY      = "SOURCE_DIR"
	RECURSE_KEY         = "RECURSE_DIR"
	FLAT_KEY            = "FLAT_OUTPUT"
	GOOGLE_CRED         = "GOOGLE_APPLICATION_CREDENTIALS" // Unused
)

var (
	template = os.Getenv(TEMPLATE_ENV_KEY)
	dest     = os.Getenv(DESTINATION_ENV_KEY)
	src      = os.Getenv(SOURCE_ENV_KEY)
	rec      = strings.ToLower(os.Getenv(RECURSE_KEY)) == "true"
	flat     = strings.ToLower(os.Getenv(FLAT_KEY)) == "true"
)

func main() {
	log.Println("🟢 Hello! Waking up to do some publishin'!")
	log.Println("🟢 Looking for Google Federated Workload Credentials.")

	// init google and connect
	google := api.GoogleApi{
		Dest:     dest,
		Template: template,
		Flat:     flat,
	}
	{
		err := google.Connect()
		if err != nil {
			log.Fatalf("🔴 Make certain you use google-github-actions/auth before this action!\n%s", err)
		} else {
			log.Println("🟢 Ambient Google credentials found! Starting grabbing files.")
		}
	}

	// init local and get paths
	// Check if src exists
	if _, err := os.Stat(src); os.IsNotExist(err) {
		log.Fatalf("🔴 Cannot find dir at given path.\n%s", err)
	}
	local := api.Local{
		Src: src,
		Rec: rec, //Conversion from environment string to actual bool
	}

	// Get paths
	var paths []string
	{
		var err error
		paths, err = local.GetPaths()
		if err != nil {
			log.Fatalf("🔴 Error retrieving paths: %s", err)
		}
	}

	log.Printf("🟢 MarkDown files found. Ready to publish %d files!", len(paths))

	{ // And now for the great dance begin!
		for _, path := range paths {
			doc, err := api.GetDoc(path)
			if err != nil {
				log.Printf("🟠 Warning: Something went wrong with retrieving the MarkDown file, skipping!\n\tpath: %s\n%s", path, err)
				continue
			}

		}
	}
}

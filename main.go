package main

// Package is called aw
import (
	"log"

	"github.com/deanishe/awgo"
)

// Workflow is the main API
var wf *aw.Workflow

func init() {
	// Create a new Workflow using default settings.
	// Critical settings are provided by Alfred via environment variables,
	// so this *will* die in flames if not run in an Alfred-like environment.
	wf = aw.New()
}

// Your workflow starts here
func run() {
	var query string

	// Use wf.Args() to enable Magic Actions
	if args := wf.Args(); len(args) > 0 {
		query = args[0]
	}
	log.Printf("[main] query=%s", query)
	// Add a "Script Filter" result
	wf.NewItem("dummy title. query: " + query).
		Subtitle("dummy subtitle").
		Arg("https://google.co.jp").
		UID("dummy uid").
		// Icon("dummy icon").
		Valid(true)

	// Send results
	wf.WarnEmpty("No matching resutl?", "Try a different query?")
	// Send results to Alfred
	wf.SendFeedback()
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}

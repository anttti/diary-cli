package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// Parse command line flags
	dir := flag.String("dir", ".", "Directory to save diary files (default: current directory)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] <text>\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Appends text to a daily diary file (YYYY-MM-DD.md)\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	// Get the text to append from remaining arguments
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: No text provided\n")
		flag.Usage()
		os.Exit(1)
	}

	text := strings.Join(args, " ")

	// Generate filename based on current date
	filename := time.Now().Format("2006-01-02") + ".md"
	filepath := filepath.Join(*dir, filename)

	// Create directory if it doesn't exist
	if err := os.MkdirAll(*dir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directory: %v\n", err)
		os.Exit(1)
	}

	// Open file in append mode (creates if doesn't exist)
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Format the entry with timestamp
	entry := fmt.Sprintf("[%s] %s\n", time.Now().Format("15:04:05"), text)

	// Write to file
	if _, err := file.WriteString(entry); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Entry added to %s\n", filepath)
}

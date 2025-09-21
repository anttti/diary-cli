package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"diary-cli/config"
)

type cliArgs struct {
	dir  string
	text string
}

func parseFlags() (*cliArgs, error) {
	defaultDir := "."
	dir := flag.String("dir", defaultDir, "Directory to save diary files (default: current directory)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] <text>\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Appends text to a daily diary file (YYYY-MM-DD.md)\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		return nil, fmt.Errorf("no text provided")
	}

	return &cliArgs{
		dir:  *dir,
		text: strings.Join(args, " "),
	}, nil
}

func setupFileAndDir(dir string, cfg *config.Config) (*os.File, error) {
	// Determine directory based on priority: CLI > Config > Default
	finalDir := config.GetDir(dir, cfg.Dir, ".")

	// Generate filename based on current date
	filename := time.Now().Format("2006-01-02") + ".md"
	filepath := filepath.Join(finalDir, filename)

	// Create directory if it doesn't exist
	if err := os.MkdirAll(finalDir, 0755); err != nil {
		return nil, fmt.Errorf("creating directory: %w", err)
	}

	// Open file in append mode (creates if doesn't exist)
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}

	return file, nil
}

func formatEntry(text string) string {
	return fmt.Sprintf("[%s] %s\n", time.Now().Format("15:04:05"), text)
}

func writeEntry(file *os.File, entry string) error {
	if _, err := file.WriteString(entry); err != nil {
		return fmt.Errorf("writing to file: %w", err)
	}
	return nil
}

func main() {
	// Load configuration from file
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to load config: %v\n", err)
		cfg = &config.Config{} // Use empty config on error
	}

	// Parse CLI flags
	args, err := parseFlags()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		flag.Usage()
		os.Exit(1)
	}

	// Setup directory and open file
	file, err := setupFileAndDir(args.dir, cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Format the entry
	entry := formatEntry(args.text)

	// Write entry to file
	if err := writeEntry(file, entry); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

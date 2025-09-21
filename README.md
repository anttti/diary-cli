# Diary CLI

A simple command-line tool for maintaining daily diary entries in markdown files.

## Installation

```bash
go build -o diary
```

## Usage

### Basic usage (saves to current directory)
```bash
./diary "Your diary entry text here"
```

### Specify a different directory
```bash
./diary -dir /path/to/diary/folder "Your diary entry text here"
```

## Features

- Creates daily markdown files with format `YYYY-MM-DD.md` (e.g., `2025-09-21.md`)
- Appends entries with timestamps in format `[HH:MM:SS]`
- Creates directories if they don't exist
- Supports custom output directories via `-dir` flag

## Examples

```bash
# Add entry to today's file in current directory
./diary "Had a great meeting today"

# Add entry to specific directory
./diary -dir ~/Documents/diary "Completed the project milestone"

# Multiple words work without quotes
./diary Started working on the new feature
```

Each entry is appended to the day's file with a timestamp:
```
[14:30:45] Had a great meeting today
[15:45:12] Completed the project milestone
[16:20:33] Started working on the new feature
```

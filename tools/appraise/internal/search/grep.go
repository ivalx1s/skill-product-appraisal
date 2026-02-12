// Package search implements scoped full-text regex search.
// Searches within a data directory for matching patterns in JSON files.
package search

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Match represents a single grep hit.
type Match struct {
	Path    string `json:"path"`    // relative to search root
	Line    int    `json:"line"`    // 1-indexed
	Content string `json:"content"` // matched line text
}

// Options controls grep behavior.
type Options struct {
	FileGlob        string
	CaseInsensitive bool
	ContextLines    int
}

// Grep searches dir recursively for lines matching pattern.
// Searches .json and .md files by default.
func Grep(dir, pattern string, opts Options) ([]Match, error) {
	if opts.CaseInsensitive {
		pattern = "(?i)" + pattern
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("invalid regex: %w", err)
	}

	var results []Match

	err = filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}

		name := d.Name()
		// Default: search .json and .md files
		if !strings.HasSuffix(name, ".json") && !strings.HasSuffix(name, ".md") {
			return nil
		}

		// Apply file glob filter
		if opts.FileGlob != "" {
			matched, _ := filepath.Match(opts.FileGlob, name)
			if !matched {
				return nil
			}
		}

		relPath, _ := filepath.Rel(dir, path)
		matches, _ := grepFile(path, relPath, re, opts.ContextLines)
		results = append(results, matches...)
		return nil
	})

	if results == nil {
		results = []Match{}
	}
	return results, err
}

func grepFile(path, relPath string, re *regexp.Regexp, contextLines int) ([]Match, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var matchNums []int
	for i, line := range lines {
		if re.MatchString(line) {
			matchNums = append(matchNums, i+1)
		}
	}
	if len(matchNums) == 0 {
		return nil, nil
	}

	if contextLines <= 0 {
		results := make([]Match, len(matchNums))
		for i, num := range matchNums {
			results[i] = Match{Path: relPath, Line: num, Content: lines[num-1]}
		}
		return results, nil
	}

	include := make(map[int]bool)
	for _, num := range matchNums {
		start := num - contextLines
		if start < 1 {
			start = 1
		}
		end := num + contextLines
		if end > len(lines) {
			end = len(lines)
		}
		for n := start; n <= end; n++ {
			include[n] = true
		}
	}

	var results []Match
	for n := 1; n <= len(lines); n++ {
		if include[n] {
			results = append(results, Match{Path: relPath, Line: n, Content: lines[n-1]})
		}
	}
	return results, nil
}

// PrintJSON outputs matches as JSON array.
func PrintJSON(matches []Match) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(matches)
}

// PrintText outputs matches in ripgrep-style format.
func PrintText(matches []Match) {
	prevPath := ""
	for _, m := range matches {
		if m.Path != prevPath {
			if prevPath != "" {
				fmt.Println()
			}
			prevPath = m.Path
		}
		fmt.Printf("%s:%d:%s\n", m.Path, m.Line, m.Content)
	}
}

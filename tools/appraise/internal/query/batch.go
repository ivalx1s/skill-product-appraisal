// Package query provides batch execution for DSL queries.
package query

import (
	"encoding/json"
	"fmt"
	"os"
)

// OutputMode controls response format.
type OutputMode int

const (
	// ModeJSON outputs standard pretty-printed JSON.
	ModeJSON OutputMode = iota
	// ModeCompact outputs minimal key:value or CSV-style format.
	ModeCompact
)

// ExecuteBatch parses and executes a full DSL query string, returning all results.
// Semicolon-separated statements are executed in order. Per-statement errors are
// captured in the results (other statements still execute).
func ExecuteBatch(queryStr string, mode OutputMode) error {
	q, err := ParseQuery(queryStr)
	if err != nil {
		return outputError(err, mode)
	}

	executor := NewExecutor()

	if len(q.Statements) == 1 {
		result, err := executor.ExecuteStatement(q.Statements[0])
		if err != nil {
			return outputError(err, mode)
		}
		return outputResult(result, mode)
	}

	// Multiple statements — return array
	results := make([]interface{}, 0, len(q.Statements))
	for _, stmt := range q.Statements {
		result, err := executor.ExecuteStatement(stmt)
		if err != nil {
			results = append(results, map[string]interface{}{
				"error": map[string]string{"message": err.Error()},
			})
			continue
		}
		results = append(results, result)
	}

	return outputResult(results, mode)
}

func outputResult(result interface{}, mode OutputMode) error {
	switch mode {
	case ModeCompact:
		return outputCompact(result)
	default:
		return outputJSON(result)
	}
}

func outputJSON(result interface{}) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(result)
}

func outputCompact(result interface{}) error {
	// For compact mode, marshal to generic representation then format
	data, err := json.Marshal(result)
	if err != nil {
		return err
	}

	// Try as map
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err == nil {
		for k, v := range m {
			fmt.Fprintf(os.Stdout, "%s:%v\n", k, formatValue(v))
		}
		return nil
	}

	// Try as array
	var arr []interface{}
	if err := json.Unmarshal(data, &arr); err == nil {
		for _, item := range arr {
			if im, ok := item.(map[string]interface{}); ok {
				for k, v := range im {
					fmt.Fprintf(os.Stdout, "%s:%v\n", k, formatValue(v))
				}
				fmt.Println("---")
			} else {
				fmt.Fprintln(os.Stdout, formatValue(item))
			}
		}
		return nil
	}

	// Fallback
	fmt.Fprintln(os.Stdout, string(data))
	return nil
}

func formatValue(v interface{}) string {
	switch tv := v.(type) {
	case string:
		return tv
	case float64:
		if tv == float64(int64(tv)) {
			return fmt.Sprintf("%d", int64(tv))
		}
		return fmt.Sprintf("%.4f", tv)
	case bool:
		if tv {
			return "true"
		}
		return "false"
	case nil:
		return ""
	default:
		b, _ := json.Marshal(v)
		return string(b)
	}
}

func outputError(err error, mode OutputMode) error {
	switch mode {
	case ModeCompact:
		fmt.Fprintf(os.Stderr, "error:%s\n", err.Error())
	default:
		enc := json.NewEncoder(os.Stderr)
		enc.SetIndent("", "  ")
		enc.Encode(map[string]interface{}{
			"error": map[string]string{"message": err.Error()},
		})
	}
	return err
}

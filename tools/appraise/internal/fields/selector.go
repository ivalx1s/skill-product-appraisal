// Package fields implements field selection and projection for DSL queries.
// Shared between DSL and any future MCP server to guarantee identical output.
package fields

import "fmt"

// ValidFields lists all recognized field names for appraisal results.
var ValidFields = map[string]bool{
	// Common result fields
	"module":         true,
	"function":       true,
	"value":          true,
	"details":        true,
	"error":          true,
	"interpretation": true,

	// BVR fields
	"bvr":              true,
	"standalone_sum":   true,
	"bundle_price":     true,
	"component_values": true,

	// Tier gap fields
	"gaps":               true,
	"from_tier":          true,
	"to_tier":            true,
	"price_gap_abs":      true,
	"price_gap_pct":      true,
	"value_gap":          true,
	"value_to_price":     true,
	"diagnosis":          true,

	// LFK fields
	"classifications": true,
	"leaders_count":   true,
	"fillers_count":   true,
	"killers_count":   true,
	"classification":  true,
	"rationale":       true,

	// Dead weight fields
	"dead_weight_ratio": true,
	"threshold":         true,
	"passes":            true,
	"dead_weight":       true,
	"component_usage":   true,

	// Financial fields
	"revenue_per_customer": true,
	"cost_per_customer":    true,
	"margin_per_customer":  true,
	"margin_pct":           true,
	"viable":               true,
	"clv":                  true,
	"break_even_units":     true,
	"fixed_costs":          true,
	"contribution_margin":  true,

	// Stress test fields
	"base_margin":     true,
	"stressed_margin": true,
	"survives_stress": true,

	// Cannibalization fields
	"net_revenue_delta": true,
	"net_positive":      true,

	// Scoring fields
	"weighted_score": true,
	"decision":       true,
	"dimensions":     true,
	"weights_used":   true,
	"risks":          true,
	"avg_score":      true,
	"max_score":      true,
	"high_risks":     true,

	// Schema fields (for introspection)
	"modules":   true,
	"functions": true,
	"schema":    true,
}

// Presets are named bundles of fields for common access patterns.
var Presets = map[string][]string{
	"minimal": {"value", "interpretation"},
	"default": {"module", "function", "value", "interpretation"},
	"full":    {"module", "function", "value", "details", "interpretation", "error"},
}

// Selector controls which fields appear in the response.
type Selector struct {
	fields map[string]bool
	all    bool
}

// NewSelector creates a Selector from requested field names.
// Empty input defaults to the "default" preset.
// Presets are expanded inline. Unknown fields return an error.
func NewSelector(requested []string) (*Selector, error) {
	if len(requested) == 0 {
		return &Selector{
			fields: map[string]bool{
				"module": true, "function": true, "value": true, "interpretation": true,
			},
		}, nil
	}

	s := &Selector{fields: make(map[string]bool)}

	for _, f := range requested {
		if expanded, ok := Presets[f]; ok {
			if f == "full" {
				s.all = true
			}
			for _, ef := range expanded {
				s.fields[ef] = true
			}
			continue
		}
		if !ValidFields[f] {
			return nil, fmt.Errorf("unknown field: %s", f)
		}
		s.fields[f] = true
	}

	return s, nil
}

// Include returns true if the field should be in the response.
func (s *Selector) Include(field string) bool {
	if s.all {
		return true
	}
	return s.fields[field]
}

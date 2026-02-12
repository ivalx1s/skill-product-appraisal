// Package fields provides projection logic for calculator results.
package fields

import (
	"encoding/json"
)

// Project takes a JSON-serializable result and returns only the fields
// selected by the Selector. Works by marshaling to a generic map,
// filtering, and returning the filtered map.
func Project(selector *Selector, result interface{}) (map[string]interface{}, error) {
	// Marshal to JSON then back to generic map
	data, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		// Result might be an array or primitive — wrap it
		return map[string]interface{}{"value": result}, nil
	}

	if selector.all {
		return m, nil
	}

	filtered := make(map[string]interface{})
	for k, v := range m {
		if selector.Include(k) {
			filtered[k] = v
		}
	}

	return filtered, nil
}

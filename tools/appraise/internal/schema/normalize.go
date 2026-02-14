package schema

import "strings"

// NormalizeQuery rewrites dotted/slashed identifiers in argument positions
// as quoted strings, so the agentquery parser (which only allows letters,
// digits, underscores, and hyphens in identifiers) can handle them.
//
// Transformation: calc(pricing.bvr, ...) → calc("pricing.bvr", ...)
//
// Only affects unquoted values that contain dots or slashes. Values already
// inside quotes are left untouched. Field names in { } projections are never
// affected (they don't contain dots).
func NormalizeQuery(input string) string {
	var b strings.Builder
	b.Grow(len(input) + 16)

	i := 0
	for i < len(input) {
		ch := input[i]

		// Pass through quoted strings untouched
		if ch == '"' {
			b.WriteByte(ch)
			i++
			for i < len(input) {
				if input[i] == '\\' && i+1 < len(input) {
					b.WriteByte(input[i])
					b.WriteByte(input[i+1])
					i += 2
					continue
				}
				if input[i] == '"' {
					b.WriteByte(input[i])
					i++
					break
				}
				b.WriteByte(input[i])
				i++
			}
			continue
		}

		// Pass through field projections { ... } untouched
		if ch == '{' {
			for i < len(input) && input[i] != '}' {
				b.WriteByte(input[i])
				i++
			}
			if i < len(input) {
				b.WriteByte(input[i])
				i++
			}
			continue
		}

		// If we hit an ident-start char, read the full "token"
		// (including dots/slashes that the old parser allowed).
		if isOldIdentStart(ch) {
			start := i
			for i < len(input) && isOldIdentChar(input[i]) {
				i++
			}
			word := input[start:i]

			// If the word contains dots or slashes, it needs quoting
			// for the agentquery parser. But only quote it if we're
			// in an argument position (after '(' or ','), not if it's
			// an operation name (before '(').
			if needsQuoting(word) && isArgPosition(input, start) {
				b.WriteByte('"')
				b.WriteString(word)
				b.WriteByte('"')
			} else {
				b.WriteString(word)
			}
			continue
		}

		b.WriteByte(ch)
		i++
	}

	return b.String()
}

// needsQuoting returns true if the word contains characters that
// the agentquery tokenizer doesn't accept in identifiers.
func needsQuoting(word string) bool {
	for _, ch := range word {
		if ch == '.' || ch == '/' {
			return true
		}
	}
	return false
}

// isArgPosition checks if the position `pos` is inside an argument list
// (i.e., preceded by '(' or ',' with optional whitespace).
func isArgPosition(input string, pos int) bool {
	// Walk backward from pos, skipping whitespace
	j := pos - 1
	for j >= 0 && (input[j] == ' ' || input[j] == '\t' || input[j] == '\n' || input[j] == '\r') {
		j--
	}
	if j < 0 {
		return false
	}
	return input[j] == '(' || input[j] == ',' || input[j] == '='
}

// isOldIdentStart mirrors the old parser's isIdentStart.
func isOldIdentStart(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '_'
}

// isOldIdentChar mirrors the old parser's isIdentChar (including dots and slashes).
func isOldIdentChar(ch byte) bool {
	return isOldIdentStart(ch) || ch == '-' || ch == '.' || ch == '/'
}

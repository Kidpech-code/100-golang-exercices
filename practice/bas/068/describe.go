package practice

// Describe returns the type name for int, string and bool,
// or "unknown" for any other type including nil.
func Describe(v any) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

// Reviewed against the exercise spec; no solution/068 in repo.

package sfmt

// Format defines how the slice should be formatted.
type Format int

// All possible format options.
const (
	Unknown Format = iota
	JSON
	YAML
	Table
	CSV
)

// ParseFormat parses a format from string.
func ParseFormat(f string) Format {
	switch f {
	case "json":
		return JSON
	case "yaml":
		return YAML
	case "table":
		return Table
	case "csv":
		return CSV
	default:
		return Unknown
	}
}

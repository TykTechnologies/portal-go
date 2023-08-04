package portal

func String(v string) *string {
	return &v
}

func StringValue(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

func Int64(v int64) *int64 {
	return &v
}

func Int64Value(s *int64) int64 {
	if s == nil {
		return 0
	}

	return *s
}

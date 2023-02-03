package validators

func UUID(s string) bool {
	if len(s) != 36 {
		return false
	}
	return true
}

package modules

func truncate(str string, length int) string {
	if len(str) < length {
		return str
	}
	return str[:length] + "…"
}

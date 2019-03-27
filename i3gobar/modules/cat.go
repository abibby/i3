package modules

func Append(cb func() string, s string) func() string {
	return func() string {
		return cb() + s
	}
}

func Prepend(s string, cb func() string) func() string {
	return func() string {
		return s + cb()
	}
}

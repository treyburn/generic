package constant

type Err string

func (ce Err) Error() string {
	return string(ce)
}

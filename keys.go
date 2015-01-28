package llrb

// IntKey is an int that satisfies the Key interface.
type IntKey int

func (k IntKey) Less(a interface{}) bool {
	return k < a.(IntKey)
}

// StringKey is an int that satisfies the Key interface.
type StringKey string

func (k StringKey) Less(a interface{}) bool {
	return k < a.(StringKey)
}

package vanda

type ErrInvalidPattern struct {
	Message string
}

func (n *ErrInvalidPattern) Error() string {
	return n.Message
}

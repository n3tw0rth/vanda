package vanda

type ArgumentParsingError struct {
	Message string
}

type ArgumentCastingError struct {
	Message string
}

func (n *ArgumentParsingError) Error() string {
	return n.Message
}

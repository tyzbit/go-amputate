package goamputate

// A request to pass to the Amputator Bot API.
type AmputationRequest struct {
	options map[string]string
	urls    []string
}

// An instance of the AmputatorBot
type AmputatorBot struct{}

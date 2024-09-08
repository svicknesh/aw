package aw

import "github.com/open-runtimes/types-for-go/v4/openruntimes"

type Header struct {
	headers map[string]string
}

// NewContextHeaders - retrieves the headers in the context for local use
func NewContextHeaders(Context openruntimes.Context) (header *Header) {
	header = new(Header)
	header.headers = make(map[string]string)
	header.headers = Context.Req.Headers
	return
}

// GetValue - returns value from a given header input
func (h *Header) GetValue(header string) (value string) {
	value, ok := h.headers[header]
	// this helps to not throw a panic if the header doesn't exist
	if !ok {
		return ""
	}

	return value
}

// GetAppwriteKey - returns Appwrite Key from the header
func (h *Header) GetAppwriteKey() (value string) {
	return h.GetValue("x-appwrite-key")
}

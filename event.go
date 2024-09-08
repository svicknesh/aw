package aw

import (
	"encoding/json"
	"strings"
)

// Event - parses the `x-appwrite-event` header to get the database, collection and document id with the action
type Event struct {
	DatabaseID   string
	CollectionID string
	DocumentID   string
	Action       string
}

// ParseEvent - parses the `x-appwrite-event` for the required details
func ParseEvent(h *Header) (e *Event) {
	e = new(Event)

	// we want the header `Context.Req.Headers["x-appwrite-event"])` since it stores the needed values, always in this format
	// databases.X.collections.Y.documents.Z.create

	// this header contains the information we are interested in
	parts := strings.Split(h.GetValue("x-appwrite-event"), ".")
	e.DatabaseID = parts[1]
	e.CollectionID = parts[3]
	e.DocumentID = parts[5]
	e.Action = parts[6]

	return
}

// String - prints the JSON string of this struct
func (e *Event) String() string {
	bytes, _ := json.Marshal(e)
	return string(bytes)
}

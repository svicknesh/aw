### Appwrite Golang Helper Library

Helper library for Appwrite. This library can do the following
- create a new client instance.
- parses the event header to get the database, collection, database id and event.
- gets header values.
- model with default values for a document.

#### Using this library

##### using custom API key
use this when writing applications in Golang

```golang
// create a new client instance using endpoint, project id and api key
client, err := aw.ClientIntsance(aw.Endpoint, os.Getenv("APPWRITE_PROJECT_ID"), os.Getenv("APPWRITE_API_KEY"))
if nil != err {
    log.Println(err.Error())
    os.Exit(1)
}

db := appwrite.NewDatabases(client)
```

##### using API key from header
use this when writing handlers for functions

```golang
// store the headers from context
h := aw.NewContextHeaders(Context)

// create a new client instance using endpoint, project id and api key
client, err := aw.ClientIntsance(aw.Endpoint, os.Getenv("APPWRITE_PROJECT_ID"), h.GetAppwriteKey())
if nil != err {
    log.Println(err.Error())
    os.Exit(1)
}

db := appwrite.NewDatabases(client)

// parse the database id, collection id, document id and action
event := aw.ParseEvent(h)

// the values from the event are stored in the following variables
// event.DatabaseID
// evente.CollectionID
// event.DocumentID
// event.Action - insert, select, update, delete

```


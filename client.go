package aw

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
)

// ClientIntsance - create new instance of Appwrite client
func ClientIntsance(endpoint, projectId, apiKey string) (client client.Client, err error) {

	if len(endpoint) == 0 {
		return client, fmt.Errorf("newclient: missing appwrite endpoint")
	}

	if len(projectId) == 0 {
		return client, fmt.Errorf("newclient: missing appwrite project ID")
	}

	if len(apiKey) == 0 {
		return client, fmt.Errorf("newclient: missing appwrite API key")
	}

	client = appwrite.NewClient(
		appwrite.WithEndpoint(endpoint),
		appwrite.WithProject(projectId),
		appwrite.WithKey(apiKey),
		//appwrite.WithKey(Context.Req.Headers["x-appwrite-key"]),
	)

	return
}

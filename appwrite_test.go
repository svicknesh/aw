package aw_test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/joho/godotenv"
	"github.com/svicknesh/aw"
)

type Colllection struct {
	aw.Model
	Name string `json:"name"`
}

func (c Colllection) String() string {
	bytes, _ := json.Marshal(c)
	return string(bytes)
}

func TestAppWrite(t *testing.T) {

	godotenv.Load(".env")

	client, err := aw.ClientIntsance(aw.Endpoint, os.Getenv("APPWRITE_PROJECT_ID"), os.Getenv("APPWRITE_API_KEY"))
	if nil != err {
		log.Println(err.Error())
		os.Exit(1)
	}

	db := appwrite.NewDatabases(client)

	response, err := db.GetDocument(os.Getenv("PROJECT_ID"), os.Getenv("COLLECTION_ID"), os.Getenv("DOCUMENT_ID"))
	if nil != err {
		return
	}

	collection := new(Colllection)
	response.Decode(collection)
	fmt.Println(collection)

}

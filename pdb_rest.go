package pdb_rest

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-resty/resty/v2"
	"golang.org/x/text/cases"
)

func initClient (debug bool, apiVersion string, baseUrl string) resty.Client {
    client := resty.New()
    client.SetRetryCount(3)

    if debug {
        client.SetDebug(true)
    }

    switch {
    case apiVersion != "" & baseUrl != "":
        client.SetBaseURL(baseUrl+apiVersion)
    case apiVersion == "" & baseUrl != "":
        client.SetBaseURL(baseUrl + "v1")
    case apiVersion != "" & baseUrl == "":
        client.SetBaseURL("https://data.rcsb.org/rest/" + apiVersion)
    default:
        client.SetBaseURL("https://data.rcsb.org/rest/v1")
}


    return client
}
    
type simpleEntryInfo  struct {
    entityID string
    entryID int
    isPolymer bool
}

func getEntriesByID (client resty.Client, entityID string) ([]simpleEntryInfo, error) {
    client.R().
    

    response, err = http.Get(fullPath)

    if err != nil {}

    
    return data.Title, err


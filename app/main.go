package main

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type jiraCreds struct {
	Username string
	Password string
	OrgName  string
}

var (
	username = flag.String("username", "", "Your Jira username")
	password = flag.String("password", "", "Your Jira password")
	org_name = flag.String("org", "Your org's name i.e. http://YOUR_ORG_NAME.atlassian.net")
	api_port = flag.String("port", "HTTP port for Gojira to listnen to")
)

func init() {
	flag.Parse()
	creds := jiraCreds{*username, *password, *jira_url}
	log.Printf("Gojira listening on port :%s", *api_port)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("api/v1/sprint", jiraSprintData).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+*api_port, router))
}

func cURLEndpoint(creds *jiraCreds, endpoint string) string {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(creds.Username, creds.Password)
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	return string(body)
}

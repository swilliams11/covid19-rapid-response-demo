// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/golang/protobuf/jsonpb"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var dfprojectID = ""
var port string
var sessionClient *dialogflow.SessionsClient
var agentClient *dialogflow.AgentsClient
var ctx = context.Background()

func main() {

	setPort()
	if err := setDialogFlowProject(); err != nil {
		log.Fatal(err)
	}

	if err := setSessionClient(); err != nil {
		log.Fatal(err)
	}
	defer sessionClient.Close()

	if err := setAgentClient(); err != nil {
		log.Fatal(err)
	}
	defer agentClient.Close()

	http.Handle("/", changeHeaderThenServe(http.FileServer(http.Dir("./dist"))))
	http.HandleFunc("/healthz", handleHealth)
	http.HandleFunc("/agent", handleTitle)
	http.HandleFunc("/query/text", handleQueryText)
	http.HandleFunc("/query/audio", handleQueryAudio)
	http.HandleFunc("/query/event", handleQueryEvent)
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// Init code
func setPort() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s\n", port)
	}
	log.Printf("Serving on port %s", port)
}

func setDialogFlowProject() error {
	var err error
	dfprojectID, err = getSecret()
	if err != nil {
		dfprojectID = os.Getenv("PROJECTDIALOGFLOW")
	}
	if dfprojectID == "" {
		return fmt.Errorf("could not get Dialogflow project from metadata or env: %v", err)
	}

	return nil
}

func getSecret() (string, error) {

	credentials, err := google.FindDefaultCredentials(ctx, compute.ComputeScope)
	if err != nil {
		return "", fmt.Errorf("could not determine this project id: %v", err)
	}
	id := credentials.ProjectID

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create secretmanager client: %v", err)
	}
	path := fmt.Sprintf("projects/%s/secrets/PROJECTDIALOGFLOW/versions/latest", id)
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: path,
	}

	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to access secret version: %v", err)
	}
	output := strings.TrimSpace(string(result.Payload.Data))

	return output, nil
}

func setSessionClient() error {
	var err error

	sessionClient, err = dialogflow.NewSessionsClient(ctx)
	if err != nil {
		return fmt.Errorf("could not get a dialogflow session client: %s", err)
	}
	return nil
}

func setAgentClient() error {
	var err error

	agentClient, err = dialogflow.NewAgentsClient(ctx)
	if err != nil {
		return fmt.Errorf("could not get a dialogflow agent client: %s", err)
	}
	return nil
}

// Handlers code.
func handleHealth(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "{\"status\":\"ok\")")
	return
}

func handleTitle(w http.ResponseWriter, r *http.Request) {
	resp, err := getAgent()
	if err != nil {
		writeError(w, err)
		return
	}

	title := resp.Description

	jsonString := fmt.Sprintf("{\"title\":\"%s\"}", title)

	writeResponse(w, http.StatusOK, jsonString)
	return
}

func getAgent() (*dialogflowpb.Agent, error) {
	req := &dialogflowpb.GetAgentRequest{
		Parent: "projects/" + dfprojectID,
	}
	resp, err := agentClient.GetAgent(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not get agent data %v", err)
	}
	return resp, nil
}

func handleQueryText(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(32 << 20)

	q := r.FormValue("q")
	sessionid := r.FormValue("session")

	resp, err := detectIntentText(dfprojectID, sessionid, q, "en")
	if err != nil {
		writeError(w, err)
		return
	}

	jsonString, err := resp.JSON()
	if err != nil {
		writeError(w, err)
		return
	}

	writeResponse(w, http.StatusOK, jsonString)
	return
}

func handleQueryAudio(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	var buf bytes.Buffer

	file, _, err := r.FormFile("file")
	if err != nil {
		nerr := fmt.Errorf("cannot get the file from the form post: %s", err)
		writeResponse(w, http.StatusInternalServerError, nerr.Error())
		return
	}

	defer file.Close()
	io.Copy(&buf, file)

	audioBytes := buf.Bytes()
	sessionid := r.FormValue("session")

	resp, err := detectIntentAudio(dfprojectID, sessionid, audioBytes, "en")
	if err != nil {
		writeError(w, err)
		return
	}

	json, err := resp.JSON()
	if err != nil {
		writeError(w, err)
		return
	}
	writeResponse(w, http.StatusOK, json)
	return
}

func handleQueryEvent(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(32 << 20)
	e := r.FormValue("event")
	sessionid := r.FormValue("session")

	resp, err := detectIntentEvent(dfprojectID, sessionid, e, "en")
	if err != nil {
		writeError(w, err)
		return
	}

	jsonString, err := resp.JSON()
	if err != nil {
		writeError(w, err)
		return
	}

	writeResponse(w, http.StatusOK, jsonString)
	return
}

// HTTP Serving related code.
func writeError(w http.ResponseWriter, err error) {
	msg := fmt.Sprintf("{\"error\":\"%s\"}", err)
	writeResponse(w, http.StatusInternalServerError, msg)
}

func writeResponse(w http.ResponseWriter, code int, msg string) {

	if code != http.StatusOK {
		log.Printf(msg)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write([]byte(msg))

	return
}

func changeHeaderThenServe(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// This page will be embedded in other pages and that causes
		// CORS issues.
		w.Header().Add("Access-Control-Allow-Origin", "*")
		// Serve with the actual handler.
		h.ServeHTTP(w, r)
	}
}

// Dialogflow related code.

// DFResponse is the structure that contains the main things we need from
// Dialogflow: the text of the original request for text or audio and
// the audio and text of the response.
type DFResponse struct {
	Text            string                         `json:"text"`
	Audio           []byte                         `json:"audio"`
	OriginalRequest string                         `json:"original_reqeust"`
	ResponseID      string                         `json:"response_id"`
	Messages        []*dialogflowpb.Intent_Message `json:"messages"`
	MesaagesJSON    []string                       `json:"messages_json"`
	Intent          *dialogflowpb.Intent           `json:"intent"`
}

// JSON Returns the given DFResponse struct as a JSON string
func (d *DFResponse) JSON() (string, error) {

	err := d.MarshalMessagesToString()
	if err != nil {
		return "", fmt.Errorf("could not marshal protobuf json for response: %s", err)
	}

	b, err := json.Marshal(d)
	if err != nil {
		return "", fmt.Errorf("could not marshal json for response: %s", err)
	}

	return string(b), nil
}

// Load just handles taking the response from Dialogflow and converting it
// for user in our application.
func (d *DFResponse) Load(response *dialogflowpb.DetectIntentResponse) {
	queryResult := response.GetQueryResult()
	d.ResponseID = response.GetResponseId()
	d.Text = queryResult.GetFulfillmentText()
	d.Audio = response.GetOutputAudio()
	d.OriginalRequest = queryResult.GetQueryText()
	d.Messages = queryResult.GetFulfillmentMessages()
	d.Intent = queryResult.GetIntent()

}

// MarshalMessagesToString is a huge kludgy mess meant to fix the problem with
// messages coming back as proto. It exports it out to raw json, and
// doesn't try to structure it, because it is not consistent therefore
// more than I wanted to tackle in this demo.
func (d *DFResponse) MarshalMessagesToString() error {

	m := jsonpb.Marshaler{}

	for _, v := range d.Messages {
		str, err := m.MarshalToString(v)
		if err != nil {
			return fmt.Errorf("could not marshal json for response: %s", err)
		}
		d.MesaagesJSON = append(d.MesaagesJSON, str)
	}

	return nil
}

func detectIntentText(dfprojectID, sessionID, text, languageCode string) (DFResponse, error) {
	result := DFResponse{}

	if dfprojectID == "" || sessionID == "" {
		return result, fmt.Errorf("received empty project or session")
	}

	sessionPath := fmt.Sprintf("projects/%s/agent/sessions/%s", dfprojectID, sessionID)
	request := dialogflowpb.DetectIntentRequest{
		Session: sessionPath,
		QueryInput: &dialogflowpb.QueryInput{
			Input: &dialogflowpb.QueryInput_Text{
				Text: &dialogflowpb.TextInput{
					Text:         text,
					LanguageCode: languageCode,
				},
			},
		},
	}

	response, err := sessionClient.DetectIntent(ctx, &request)
	if err != nil {
		return result, err
	}

	result.Load(response)

	return result, nil
}

func detectIntentEvent(dfprojectID, sessionID, eventName, languageCode string) (DFResponse, error) {
	result := DFResponse{}

	if dfprojectID == "" || sessionID == "" {
		return result, fmt.Errorf("received empty project or session")
	}

	sessionPath := fmt.Sprintf("projects/%s/agent/sessions/%s", dfprojectID, sessionID)
	request := dialogflowpb.DetectIntentRequest{
		Session: sessionPath,
		QueryInput: &dialogflowpb.QueryInput{
			Input: &dialogflowpb.QueryInput_Event{
				Event: &dialogflowpb.EventInput{
					Name:         eventName,
					LanguageCode: languageCode,
				},
			},
		},
	}

	response, err := sessionClient.DetectIntent(ctx, &request)
	if err != nil {
		return result, err
	}

	result.Load(response)
	return result, nil
}

func detectIntentAudio(dfprojectID, sessionID string, audioBytes []byte, languageCode string) (DFResponse, error) {

	result := DFResponse{}

	if dfprojectID == "" || sessionID == "" {
		return result, fmt.Errorf("Received empty project or session")
	}

	sessionPath := fmt.Sprintf("projects/%s/agent/sessions/%s", dfprojectID, sessionID)

	// In this example, we hard code the encoding and sample rate for simplicity.
	request := dialogflowpb.DetectIntentRequest{
		Session:    sessionPath,
		InputAudio: audioBytes,
		QueryInput: &dialogflowpb.QueryInput{
			Input: &dialogflowpb.QueryInput_AudioConfig{
				AudioConfig: &dialogflowpb.InputAudioConfig{
					LanguageCode: languageCode,
				},
			},
		},
	}

	response, err := sessionClient.DetectIntent(ctx, &request)
	if err != nil {
		return result, err
	}

	result.Load(response)

	return result, nil
}

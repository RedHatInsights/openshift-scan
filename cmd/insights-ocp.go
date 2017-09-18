package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var databaseDir = ".insights-ocp-data"

type Request struct {
	ContentPath string
	ImageId     string
}

type System struct {
	Hostname string            `json:"hostname"`
	Metadata map[string]string `json:"metadata"`
}

type Content struct {
	Plain string `json:"plain"`
	Html  string `json:"html"`
}

type Report struct {
	RuleData       map[string]interface{} `json:"rule_data"`
	Title          *Content               `json:"title"`
	Summary        *Content               `json:"summary"`
	Description    *Content               `json:"description"`
	Details        *Content               `json:"details"`
	Reference      *Content               `json:"reference"`
	Resolution     *Content               `json:"resolution"`
	Severity       string                 `json:"severity"`
	Category       string                 `json:"category"`
	Impact         int                    `json:"impact"`
	Likelihood     int                    `json:"likelihood"`
	RebootRequired bool                   `json:"reboot_required"`
	Acks           []interface{}          `json:"acks"`
}

type ClientResponse struct {
	Version string                 `json:"version"`
	System  *System                `json:"system"`
	Reports map[string]Report      `json:"reports"`
	Upload  map[string]interface{} `json:"upload"`
}

// TODO can we just depend on this code from image inspector directly?
type ImageInspectorResponse struct {
	Name           string                  `json:"name"`
	ScannerVersion string                  `json:"scannerVersion"`
	TimeStamp      time.Time               `json:"timestamp"`
	Reference      string                  `json:"reference"`
	Description    string                  `json:"description"`
	Summary        []ImageInspectorSummary `json:"summary`
}

type ImageInspectorSummary struct {
	Label Severity `json:"label"`
}

type Severity string

func inspect(w http.ResponseWriter, r *http.Request) {
	var req Request
	if r.Body == nil {
		http.Error(w, "Please send request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if req.ContentPath == "" || req.ImageId == "" {
		http.Error(w, "Invalid request", 400)
		return
	}
	//TODO switch to log
	fmt.Println(req.ContentPath)
	fmt.Println(req.ImageId)
	result, err := scanImage(req.ContentPath, req.ImageId)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Unable to scan image.", 500)
		return
	} else {
		emit(w, result)
	}
}

func persist(imageId string, data *ClientResponse) {
	var targ = filepath.Join(databaseDir, imageId[:2], imageId) + ".json"
	os.MkdirAll(filepath.Dir(targ), 0700)
	fp, _ := os.Create(targ)
	json_text, _ := json.Marshal(&data)
	fp.Write(json_text)
	fp.Close()
}

func echo(w http.ResponseWriter, r *http.Request) {
	var doc ClientResponse
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	emit(w, &doc)
}

func emit(w http.ResponseWriter, resp *ClientResponse) {
	w.Header().Set("Content-Type", "application/json")
	var json_text, _ = json.Marshal(resp)
	w.Write(json_text)
}

func scanImage(contentPath string, imageId string) (*ClientResponse, error) {
	cmdStr := "insights-client --analyze-image --verbose --to-json --mountpoint " + contentPath
	out, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()
	fmt.Printf("%s", cmdStr+"\n")
	fmt.Printf("%s", out)
	var res ClientResponse
	json.Unmarshal(out, &res)
	persist(imageId, &res)
	return &res, nil
}

func main() {
	http.HandleFunc("/inspect", inspect)
	http.HandleFunc("/echo", echo)
	http.ListenAndServe(":9000", nil) //TODO don't hard code port
}

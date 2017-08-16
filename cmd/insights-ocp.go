package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "os/exec"
)

type Request struct{
    ContentPath      string
    ImageId   string
}

type Response struct{
    Body []byte
}

func inspect(w http.ResponseWriter, r *http.Request) {
    var req Request
    if r.Body == nil {
       http.Error(w,"Please send request body", 400)
       return
    }
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
            http.Error(w, err.Error(), 400)
            return
    }
    fmt.Println(req.ContentPath)
    fmt.Println(req.ImageId)
    result, err := scanImage(req.ContentPath,req.ImageId)
    w.Header().Set("Content-Type", "application/json")
    w.Write(result.Body)
}


func scanImage(contentPath string, imageId string) (*Response,error){
    
    cmdStr := "insights-client --mountpoint " + contentPath
    out, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()  
    fmt.Printf("%s", cmdStr+ "\n")
    fmt.Printf("%s", out)
    body,err := loadResults()
    if err != nil {
        return nil, err
    }
    return &Response{Body: body},nil
}



func loadResults() ([]byte, error) {
    body, err := ioutil.ReadFile("/etc/insights-client/.last-upload.results")
    if err != nil {
        return nil, err
    }
    return body,nil
}

func main() {
    http.HandleFunc("/inspect", inspect)
    http.ListenAndServe(":9000", nil)
}


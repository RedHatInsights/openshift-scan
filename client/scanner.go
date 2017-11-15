package client

import (
	"encoding/json"
	"os/exec"
	"fmt"


	"github.com/RedHatInsights/insights-goapi/common"
)

const clientCmd = "sudo insights-client --no-gpg --analyze-mountpoint=%s"

type Scanner interface {
	
	ScanImage(contentPath string, imageId string) (*common.ScanResponse, *[]byte, error)

}

type DefaultScanner struct{}

func NewDefaultScanner()(Scanner){
	return &DefaultScanner{}
}

func (s * DefaultScanner )ScanImage(contentPath string, imageId string) (response *common.ScanResponse, rawResponse *[]byte, err error) {
	cmdStr := fmt.Sprintf(clientCmd, contentPath)
	*rawResponse, err = exec.Command("/bin/sh", "-c", cmdStr).Output()
    if err != nil {
		return 
	}
	err = json.Unmarshal(*rawResponse, &response)
	return 
}
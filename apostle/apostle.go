package apostle

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	conf = loadConfig()
)

func Send(templateId string, email string, name string, data map[string]string) error {
	// generate json request body
	m := map[string]interface{}{
		"recipients": map[string]interface{}{
			email: map[string]interface{}{
				"template_id": templateId,
				"name":        name,
				"data":        data,
			},
		},
	}
	requestJson, err := json.Marshal(m)
	if err != nil {
		return err
	}
	log.Printf("Sending Apostle: %v", string(requestJson))
	postData := bytes.NewBuffer(requestJson)

	// prepare request
	req, err := http.NewRequest("POST", "http://deliver.apostle.io", postData)

	// set appropriate headers for auth/content type etc
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("APOSTLE_BEARER_TOKEN")))

	// do request
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// apostle returns 202 Accepted if it's gonna send
	if response.StatusCode != http.StatusAccepted {
		return errors.New(fmt.Sprintf("HTTP error from apostle of code %v", response.StatusCode))
	}
	readBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print(err)
	}
	log.Printf("Apostle returned: %v", string(readBody))

	return err
}

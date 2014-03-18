package apostle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Queue struct {
	Recipients map[string]Mail `json:"recipients"`
}

func NewQueue() Queue {
	return Queue{make(map[string]Mail)}
}

func (q *Queue) Add(m Mail) (length int, err error) {
	if len(m.Email) == 0 {
		err = NoEmailError{}
		return
	}
	q.Recipients[m.Email] = m
	return q.Size(), nil
}

func (q *Queue) Size() int {
	return len(q.Recipients)
}

func (q *Queue) Deliver() (err error) {
	// Apostle must be configured with a domain key
	if len(conf.DomainKey) == 0 {
		return fmt.Errorf("No DomainKey is set. Provide one via ENV['APOSTLE_DOMAIN_KEY'], or call apostle.SetDomainKey()")
	}

	payload, err := json.Marshal(q)
	if err != nil {
		return
	}

	data := bytes.NewBuffer(payload)
	req, err := http.NewRequest("POST", conf.DeliveryHost, data)
	if err != nil {
		return
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", conf.DomainKey))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusAccepted {
		if resp.StatusCode == 403 {
			return InvalidDomainKeyError{req, resp}
		}
		if resp.StatusCode == 422 {
			return InvalidDomainKeyError{req, resp}
		}
		if resp.StatusCode >= 500 && resp.StatusCode < 600 {
			return ServerError{req, resp}
		}
		return DeliveryError{req, resp}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	log.Printf("Apostle returned: %v", string(body))

	return
}

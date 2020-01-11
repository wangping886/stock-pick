package httpclient

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

//wrapper
type Httpclient struct {
	*http.Client
}

func DefaultHttpclient() *Httpclient {
	return &Httpclient{http.DefaultClient}
}

func (client *Httpclient) HttpGet(url string) (data string, err error) {
	start := time.Now()
	defer func() {
		log.Println(url, data, time.Since(start), err)
	}()
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err = getStrFromResp(resp)
	if err != nil {
		return "", err
	}
	return data, nil
}

func getStrFromResp(resp *http.Response) (string, error) {
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	str := strings.Trim(string(data), "\r\n")
	if resp.StatusCode != http.StatusOK {
		return str, errors.New(fmt.Sprintf("statusCode=%v", resp.StatusCode))
	}
	return str, nil
}

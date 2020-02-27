package ip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"unblocker/model"
)

type IP struct{}

func New() *IP {
	return &IP{}
}

func (ip *IP) GetListIP(url string) (model.Response, error) {
	var (
		modelIP model.Response
	)

	url = fmt.Sprintf("https://dns.google.com/resolve?name=%s&type=A", url)
	resp, err := http.Get(url)
	if err != nil {
		return modelIP, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return modelIP, err
	}
	err = json.Unmarshal(body, &modelIP)
	if err != nil {
		return modelIP, err
	}

	return modelIP, nil
}

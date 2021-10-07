package main

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:20000/home", strings.NewReader(""))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("body=%v\n", string(bs))
	for k, v := range resp.Header {
		logrus.Infof("key=%v|value=%v\n", k, v)
	}
}

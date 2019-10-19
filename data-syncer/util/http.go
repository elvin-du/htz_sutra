package util

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func HTTPGetJson(url string, v interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if nil != err {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		return err
	}
	defer resp.Body.Close()

	bin, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return err
	}

	if http.StatusOK != resp.StatusCode {
		return NewHTTPError(resp.StatusCode, string(bin))
	}
	log.Debugln("httpget body:", string(bin))

	err = json.Unmarshal(bin, v)
	if nil != err {
		return err
	}

	return nil
}

func HTTPGetString(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if nil != err {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		return "", err

	}
	defer resp.Body.Close()

	bin, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return "", err
	}

	if http.StatusOK != resp.StatusCode {
		return "", NewHTTPError(resp.StatusCode, string(bin))
	}
	log.Debugln("httpget body:", string(bin))

	return string(bin), nil
}

type HTTPError struct {
	Code int
	Msg  string
}

func NewHTTPError(code int, msg string) *HTTPError {
	return &HTTPError{code, msg}
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf(`{"code":%d,"msg":"%s"}`, e.Code, e.Msg)
}

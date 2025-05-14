package service

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

type VshepServiceImpl struct {
	vshepURL string
}

func NewVshepServiceImpl(vshepURL string) *VshepServiceImpl {
	return &VshepServiceImpl{vshepURL}
}

func (v *VshepServiceImpl) SendRequest(request []byte) (int, []byte, error) {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100
	client := &http.Client{
		Timeout:   45 * time.Second,
		Transport: t,
	}
	req, err := http.NewRequest(http.MethodPost, v.vshepURL, bytes.NewBuffer(request))
	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("Error in calling vshep: %v", err)
		return 500, nil, err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("Error in calling vshep: %v", err)
		return 500, nil, err
	}
	logrus.Info("Vshep Status: " + resp.Status)
	logrus.Info("Vshep Header: ")
	logrus.Info(resp.Header)
	logrus.Info("Vshep Body: ")
	logrus.Info(string(respBytes))

	return resp.StatusCode, respBytes, nil
}

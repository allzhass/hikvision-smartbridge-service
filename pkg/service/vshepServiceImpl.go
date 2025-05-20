package service

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"regexp"
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

	logrus.WithFields(logrus.Fields{
		"message_id": getMessageId(request),
		"status":     resp.StatusCode,
		"response":   string(respBytes),
	}).Info("Vshep")
	return resp.StatusCode, respBytes, nil
}

func getMessageId(request []byte) string {
	re := regexp.MustCompile(`<message_id>(.*?)</message_id>`)
	match := re.FindStringSubmatch(string(request))
	if len(match) > 1 {
		return match[1]
	} else {
		return "error: message_id not found"
	}
}

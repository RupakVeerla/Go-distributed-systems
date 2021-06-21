package log

import (
	"bytes"
	"distributedServices/registry"
	"fmt"
	stlog "log"
	"net/http"
)

type clientLogger struct {
	url string
}

func SetClientLogger(serviceURL string, clientServices registry.ServiceName) {
	stlog.SetPrefix(fmt.Sprintf("[%v] - ", clientServices))
	stlog.SetFlags(0)
	stlog.SetOutput(&clientLogger{url: serviceURL})
}

func (cl clientLogger) Write(data []byte) (int, error) {
	b := bytes.NewBuffer([]byte(data))
	res, err := http.Post(cl.url+"/log", "text/plain", b)
	if err != nil {
		return 0, err
	}
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Failed to send log message. Service responded with status code: %v\n", res.StatusCode)
	}
	return len(data), nil
}

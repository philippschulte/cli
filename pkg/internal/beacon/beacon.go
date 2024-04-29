package beacon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fastly/cli/pkg/api/undocumented"
	"github.com/fastly/cli/pkg/global"
)

const (
	StatusSuccess = "success"
	StatusFail    = "fail"
)

type Event struct {
	Name    string         `json:"event"`
	Status  string         `json:"status"`
	Payload map[string]any `json:"payload"`
}

const beaconNotify = "/cli/%s/notify"

func Notify(g *global.Data, serviceID string, e Event) error {
	headers := []undocumented.HTTPHeader{
		{
			Key:   "Content-Type",
			Value: "application/json",
		},
	}

	body, err := json.Marshal(e)
	if err != nil {
		return err
	}

	co := undocumented.CallOptions{
		APIEndpoint: "https://fastly-notification-relay.edgecompute.app",
		Path:        fmt.Sprintf(beaconNotify, serviceID),
		Method:      http.MethodPost,
		HTTPHeaders: headers,
		HTTPClient:  g.HTTPClient,
		Body:        bytes.NewReader(body),
	}

	_, err = undocumented.Call(co)
	if err != nil {
		return err
	}

	return nil
}

package pushwoosh

import (
	"context"
	"net/http"
	"strconv"
	"strings"
)

type GetMsgStatsResult struct {
	StatusCode    int64  `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Repsonse      struct {
		RequestId string `json:"request_id"`
	} `json:"response"`
}

// GetMsgStats starts push statistics receiving.
// https://www.pushwoosh.com/reference/#getmsgstats
func (c *Client) GetMsgStats(ctx context.Context, message string) (*GetMsgStatsResult, error) {
	var result GetMsgStatsResult
	msg := map[string]interface{}{
		"message": message,
	}

	err := c.call(ctx, http.MethodPost, "getMsgStats", msg, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type GetResultsRepsonse struct {
	Formatter string          `json:"formatter"`
	Rows      []GetResultsRow `json:"rows"`
}

type GetResultsRow struct {
	Datetime string `json:"datetime"`
	Action   string `json:"action"`
	Count    Count  `json:"count"`
}

type Count int64

func (c *Count) UnmarshalJSON(data []byte) error {
	if string(data) == "0" {
		*c = 0
	} else {
		n, err := strconv.ParseInt(strings.Trim(string(data), `"`), 10, 64)
		if err != nil {
			return err
		}

		*c = Count(n)
	}

	return nil
}

type GetResultsResult struct {
	StatusCode    int64              `json:"status_code"`
	StatusMessage string             `json:"status_message"`
	Repsonse      GetResultsRepsonse `json:"response"`
}

// GetResults returns result of push statistics receiving.
// http://docs.pushwoosh.com/docs/createmessage
func (c *Client) GetResults(ctx context.Context, requestId string) (*GetResultsResult, error) {
	var result GetResultsResult
	notifications := map[string]interface{}{
		"request_id": requestId,
	}

	err := c.call(ctx, http.MethodPost, "getResults", notifications, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

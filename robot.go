package drobot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type DingTalkRobot struct {
	Webhook string
	Secret  string
}

type dingTalkResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (d *DingTalkRobot) SendMarkdown(title string, content string) error {
	message := "{\"msgtype\": \"markdown\",\"markdown\": {\"title\":\"" + title + "\",\"text\": \"" + content + "\"}}"

	timestamp := time.Now().UnixMilli()
	string2sign := fmt.Sprintf("%d\n%s", timestamp, d.Secret)
	sign := url.QueryEscape(hmacSHA256(string2sign, d.Secret))

	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("%s&timestamp=%d&sign=%s", d.Webhook, timestamp, sign)
	req, err := http.NewRequest("POST", url, strings.NewReader(message))
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var result dingTalkResponse
	json.Unmarshal(data, &result)
	if result.Errcode != 0 {
		return fmt.Errorf("dingding response error, code: %d, message: %s", result.Errcode, result.Errmsg)
	}
	return nil
}

func hmacSHA256(plainText string, screct string) string {
	key := []byte(screct)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(plainText))
	cipherText := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(cipherText)
}

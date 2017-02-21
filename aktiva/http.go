package aktiva

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

const epSendInvoice = "sendinvoice"
const epSendPurchaseInvoice = "sendpurchinvoice"
const epSendGeneralLedger = "sendglbatch"
const epGetTaxes = "gettaxes"

func (a *Aktiva) httpGet(ep string, resp interface{}) error {
	var err error

	timestamp := TimeToString(time.Now())
	signature := a.getSignature(timestamp, []byte(""))

	url := a.apiUrl + ep + "?ApiId=" + a.apiId + "&timestamp=" + timestamp + "&signature=" + signature

	hresp, err := http.Get(url)
	if err != nil {
		return err
	}

	if hresp.StatusCode != 200 {
		var msg struct {
			Message string
		}
		if err = json.NewDecoder(hresp.Body).Decode(&msg); err != nil {
			return err
		}
		return errors.New("HTTP Status: " + hresp.Status + "  Message: " + msg.Message)
	}

	var t string

	if resp != nil {
		if err = json.NewDecoder(hresp.Body).Decode(&t); err != nil {
			return err
		}
		if err = json.Unmarshal([]byte(t), resp); err != nil {
			return err
		}
	}

	return nil
}

func (a *Aktiva) httpPost(ep string, req interface{}, resp interface{}) error {
	var err error

	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(req)
	if err != nil {
		return err
	}

	timestamp := TimeToString(time.Now())
	signature := a.getSignature(timestamp, b.Bytes())

	url := a.apiUrl + ep + "?ApiId=" + a.apiId + "&timestamp=" + timestamp + "&signature=" + signature

	hresp, err := http.Post(url, "application/json", b)
	if err != nil {
		return err
	}

	var t string
	if hresp.StatusCode != 200 {
		var msg struct {
			Message string
		}
		if err = json.NewDecoder(hresp.Body).Decode(&msg); err != nil {
			return err
		}
		return errors.New("HTTP Status: " + hresp.Status + "  Message: " + msg.Message)
	}

	if resp != nil {
		if err = json.NewDecoder(hresp.Body).Decode(&t); err != nil {
			return err
		}
		if err = json.Unmarshal([]byte(t), resp); err != nil {
			return err
		}
	}

	return nil
}

func (a *Aktiva) getSignature(timestamp string, payload []byte) string {
	mac := hmac.New(sha256.New, []byte(a.apiKey))

	mac.Write([]byte(a.apiId))
	mac.Write([]byte(timestamp))
	mac.Write(payload)

	return base64.URLEncoding.EncodeToString(mac.Sum(nil))
}

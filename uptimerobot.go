package uptimerobot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var NotImplementedErr = fmt.Errorf("Not implemented")

type UptimeRobot struct {
	apiKey string
	client http.Client
}

func New(apiKey string) *UptimeRobot {
	return &UptimeRobot{
		apiKey: apiKey,
		client: http.Client{},
	}
}

func (r *UptimeRobot) request(method string, reqUrl *url.URL, body io.Reader) (resp *http.Response, err error) {

	reqUrl.Scheme = "https"
	reqUrl.Host = "api.uptimerobot.com"

	q := reqUrl.Query()
	q.Set("apiKey", r.apiKey)
	q.Set("format", "json")
	q.Set("noJsonCallback", "1")

	reqUrl.RawQuery = q.Encode()

	req, err := http.NewRequest(method, reqUrl.String(), body)

	if err != nil {
		return
	}

	return r.client.Do(req)
}

func (r *UptimeRobot) GetMonitors() (m []Monitor, err error) {
	reqUrl := &url.URL{
		Path: "getMonitors",
	}

	resp, err := r.request("GET", reqUrl, nil)
	if err != nil {
		return
	}

	jsonResponse := new(GetMonitors)

	if err = json.NewDecoder(resp.Body).Decode(jsonResponse); err != nil {
		return
	}

	m = jsonResponse.Monitors.Monitors
	return
}

func (r *UptimeRobot) GetAccountDetails() (a Account, err error) {
	reqUrl := &url.URL{
		Path: "getAccountDetails",
	}

	resp, err := r.request("GET", reqUrl, nil)
	if err != nil {
		return
	}

	jsonResponse := new(GetAccountDetails)

	if err = json.NewDecoder(resp.Body).Decode(jsonResponse); err != nil {
		return
	}

	a = jsonResponse.Account
	return
}

func (r *UptimeRobot) NewMonitor() (id int, err error) {
	err = NotImplementedErr
	return
}

func (r *UptimeRobot) EditMonitor() (id int, err error) {
	err = NotImplementedErr
	return
}

func (r *UptimeRobot) DeleteMonitor() (id int, err error) {
	err = NotImplementedErr
	return
}

func (r *UptimeRobot) ResetMonitor() (id int, err error) {
	err = NotImplementedErr
	return
}

func (r *UptimeRobot) GetAlertContacts() (c []AlertContact, err error) {
	reqUrl := &url.URL{
		Path: "getAlertContacts",
	}

	resp, err := r.request("GET", reqUrl, nil)
	if err != nil {
		return
	}

	jsonResponse := new(GetAlertContacts)

	if err = json.NewDecoder(resp.Body).Decode(jsonResponse); err != nil {
		return
	}

	c = jsonResponse.AlertContacts.AlertContact
	return
}
func (r *UptimeRobot) NewAlertContact() (id int, err error) {
	err = NotImplementedErr
	return
}
func (r *UptimeRobot) DeleteAlertContact() (id int, err error) {
	err = NotImplementedErr
	return
}

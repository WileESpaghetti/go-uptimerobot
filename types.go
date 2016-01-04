package uptimerobot

import "strconv"

type MonitorType int
type MonitorSubtype int
type MonitorStatus int
type MonitorKeywordType int
type Port int
type ContactType int
type ContactStatus int

func unmarshalMaybeInt(b []byte) (val int, err error) {
	strVal := string(b[1 : len(b)-1])
	if strVal == "" {
		return
	}
	val64, err := strconv.ParseInt(strVal, 10, 32)
	val = int(val64)
	return
}

func (t *MonitorType) UnmarshalJSON(b []byte) (err error) {
	val, err := unmarshalMaybeInt(b)
	if err == nil {
		*t = MonitorType(val)
	}
	return
}

func (t *MonitorSubtype) UnmarshalJSON(b []byte) (err error) {
	val, err := unmarshalMaybeInt(b)
	if err == nil {
		*t = MonitorSubtype(val)
	}
	return
}

func (t *MonitorStatus) UnmarshalJSON(b []byte) (err error) {
	val, err := unmarshalMaybeInt(b)
	if err == nil {
		*t = MonitorStatus(val)
	}
	return
}

func (t *MonitorKeywordType) UnmarshalJSON(b []byte) (err error) {
	val, err := unmarshalMaybeInt(b)
	if err == nil {
		*t = MonitorKeywordType(val)
	}
	return
}

func (t *Port) UnmarshalJSON(b []byte) (err error) {
	val, err := unmarshalMaybeInt(b)
	if err == nil {
		*t = Port(val)
	}
	return
}

func (t *ContactType) UnmarshalJSON(b []byte) (err error) {
	val, err := unmarshalMaybeInt(b)
	if err == nil {
		*t = ContactType(val)
	}
	return
}

func (t *ContactStatus) UnmarshalJSON(b []byte) (err error) {
	val, err := unmarshalMaybeInt(b)
	if err == nil {
		*t = ContactStatus(val)
	}
	return
}

type GetMonitors struct {
	Stat     string   `json:"stat"`
	Offset   int      `json:"offset,string"`
	Limit    int      `json:"limit,string"`
	Total    int      `json:"total,string"`
	Monitors Monitors `json:"monitors"`
}

type Monitors struct {
	Monitors []Monitor `json:"monitor"`
}

type Monitor struct {
	Id                int                `json:"id,string"`
	FriendlyName      string             `json:"friendlyname"`
	Url               string             `json:"url"`
	Type              MonitorType        `json:"type"`
	SubType           MonitorSubtype     `json:"subtype"`
	KeywordType       MonitorKeywordType `json:"keywordtype"`
	KeywordValue      string             `json:"keywordvalue"`
	HttpUsername      string             `json:"httpusername"`
	HttpPassword      string             `json:"httppassword"`
	Port              Port               `json:"port"`
	Interval          int                `json:"interval,string"`
	Status            MonitorStatus      `json:"status"`
	AllUptimeRatio    float32            `json:"alluptimeratio,string"`
	CustomUptimeRatio float32            `json:"customuptimeratio,string"`
}

type GetAccountDetails struct {
	Stat    string  `json:"stat"`
	Offset  int     `json:"offset"`
	Limit   int     `json:"limit"`
	Total   int     `json:"total"`
	Account Account `json:"account"`
}

type Account struct {
	MonitorLimit    int `json:"monitorLimit,string"`
	MonitorInterval int `json:"monitorInterval,string"`
	UpMonitors      int `json:"upMonitors,string"`
	DownMonitors    int `json:"downMonitors,string"`
	PausedMonitors  int `json:"pausedMonitors,string"`
}

type GetAlertContacts struct {
	Stat          string `json:"stat"`
	Offset        int    `json:"offset,string"`
	Limit         int    `json:"limit,string"`
	Total         int    `json:"total,string"`
	AlertContacts AlertContacts
}

type AlertContacts struct {
	AlertContact []AlertContact
}

type AlertContact struct {
	Id           int           `json:"id,string"`
	Type         ContactType   `json:"type"`
	Value        string        `json:"value"`
	FriendlyName string        `json:"friendlyname"`
	Status       ContactStatus `json:"status"`
	Treshold     string        `json:"treshold"`
	Recurrence   string        `json:"recurrence"`
}

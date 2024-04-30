package conf

import (
	"encoding/json"
	"reflect"
	"regexp"
)

// Source is a source configuration.
type Group struct {
	Regexp       *regexp.Regexp `json:"-"`    // filled by Check()
	Name         string         `json:"name"` // filled by Check()
	Enabled      bool           `json:"enabled"`
	EnablePublic bool           `json:"enable_public"`
	AllowPublish bool           `json:"allow_publish"`
	AllowRTSP    bool           `json:"allow_rtsp"`
	AllowWebRTC  bool           `json:"allow_webrtc"`
	AllowHLS     bool           `json:"allow_hls"`

	// General
	Paths map[string]*Path `json:"paths"`

	// Authentication
	PublishUser Credential `json:"publishUser"`
	PublishPass Credential `json:"publishPass"`
	PublishIPs  IPsOrCIDRs `json:"publishIPs"`
	ReadUser    Credential `json:"readUser"`
	ReadPass    Credential `json:"readPass"`
	ReadIPs     IPsOrCIDRs `json:"readIPs"`

	// Publisher source
	OverridePublisher        bool   `json:"overridePublisher"`
	DisablePublisherOverride *bool  `json:"disablePublisherOverride,omitempty"` // deprecated
	SRTPublishPassphrase     string `json:"srtPublishPassphrase"`

	// Hooks
	RunOnInit                  string         `json:"runOnInit"`
	RunOnInitRestart           bool           `json:"runOnInitRestart"`
	RunOnDemand                string         `json:"runOnDemand"`
	RunOnDemandRestart         bool           `json:"runOnDemandRestart"`
	RunOnDemandStartTimeout    StringDuration `json:"runOnDemandStartTimeout"`
	RunOnDemandCloseAfter      StringDuration `json:"runOnDemandCloseAfter"`
	RunOnUnDemand              string         `json:"runOnUnDemand"`
	RunOnReady                 string         `json:"runOnReady"`
	RunOnReadyRestart          bool           `json:"runOnReadyRestart"`
	RunOnNotReady              string         `json:"runOnNotReady"`
	RunOnRead                  string         `json:"runOnRead"`
	RunOnReadRestart           bool           `json:"runOnReadRestart"`
	RunOnUnread                string         `json:"runOnUnread"`
	RunOnRecordSegmentCreate   string         `json:"runOnRecordSegmentCreate"`
	RunOnRecordSegmentComplete string         `json:"runOnRecordSegmentComplete"`
}

// Equal checks whether two Paths are equal.
func (pconf *Group) Equal(other *Group) bool {
	return reflect.DeepEqual(pconf, other)
}

// Clone clones the configuration.
func (pconf Group) Clone() *Group {
	enc, err := json.Marshal(pconf)
	if err != nil {
		panic(err)
	}

	var dest Group
	err = json.Unmarshal(enc, &dest)
	if err != nil {
		panic(err)
	}

	dest.Regexp = pconf.Regexp

	return &dest
}

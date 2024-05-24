package conf

type Snapshot struct {
	Source   string `json:"source"`
	Enabled  bool   `json:"enabled"`
	Interval uint32 `json:"interval"`
	OnDemand bool   `json:"onDemand"`
	Width    uint16 `json:"width"`
	Height   uint16 `json:"height"`
}

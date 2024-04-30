package conf

type Transcoder struct {
	Id       string `json:"id"`
	Enabled  bool   `json:"enabled"`
	SourceId string `json:"source_id"`
	//Width     int    `json:"width"`
	//Height    int    `json:"height"`
	Framerate int    `json:"frame_rate"`
	Codec     string `json:"codec"`
	Bitrate   int64  `json:"bitrate"`
	GOP       int    `json:"gop"`
	Scale     struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"scale", omitempty`
	Overlay struct {
		Image      string `json:"image"`
		OffsetTop  int    `json:"offset_top"`
		OffsetLeft int    `json:"offset_left"`
	} `json:"overlay", omitempty`
}

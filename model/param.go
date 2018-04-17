package model

import (
	"encoding/json"
	"errors"
	"time"
)

// DDSParam is the DDS mode in which a DDS controllable parameter runs.
//
// Note that usually only one of the embedded structs will be not nil as it
// does not make sense to have a constant frequency which we is also swept.
type DDSParam struct {
	*DDSConst
	*DDSSweep
	*DDSPlayback
}

// DDSConst is the DDS mode where a DDS controllable parameter is constant.
type DDSConst struct {
	Value float64 `validation:"required" json:"value"`
}

// DDSSweep is the DDS mode where a DDS controllable parameter is swept.
type DDSSweep struct {
	Limits   [2]float64    `validation:"len=2,range,dive,gt=0" json:"limits"`
	NoDwell  [2]bool       `validation:"len=2"                 json:"nodwell"`
	Duration time.Duration `validation:"required,gt=0"         json:"duration"`
}

// DDSPlayback is the DDS mode where a DDS controllable parameter is playbed
// back from memory.
type DDSPlayback struct {
	WithTrigger bool    `                  json:"trigger"`
	WithDuplex  bool    `                  json:"duplex"`
	Rate        float64 `validation:"gt=0" json:"rate"`
	Data        []byte  `                  json:"data"`
}

// DDSParamJSON is a helper to decode json into the corresponding DDS mode.
//
// Because it is only valid to have one of the previous defined DDS modes
// enabled we will have on json object assigned to a controllable parameter
// field and decode it into the corresponding DDS mode struct. To detect what
// DDS mode struct is correct we introduce an additional mode field which
// holds the name of the mode i.e. "const", "sweep", "playback". This allows us
// to distinct invalid mode errors from decoding errors.
type DDSParamJSON struct {
	Mode string `json:"mode"`
}

// MarshalJSON implements json.Marshaler interface.
func (d DDSParam) MarshalJSON() ([]byte, error) {
	if d.DDSConst != nil {
		return json.Marshal(d.DDSConst)
	}
	if d.DDSSweep != nil {
		return json.Marshal(d.DDSSweep)
	}
	if d.DDSPlayback != nil {
		return json.Marshal(d.DDSPlayback)
	}

	return json.Marshal(nil)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (d *DDSParam) UnmarshalJSON(b []byte) error {
	p := &DDSParamJSON{}

	if err := json.Unmarshal(b, p); err != nil {
		return err
	}

	switch p.Mode {
	case "const":
		d.DDSConst = &DDSConst{}

		return json.Unmarshal(b, d.DDSConst)
	case "sweep":
		d.DDSSweep = &DDSSweep{}

		return json.Unmarshal(b, d.DDSSweep)
	case "playback":
		d.DDSPlayback = &DDSPlayback{}

		return json.Unmarshal(b, d.DDSPlayback)
	}

	return errors.New("unsupported mode")
}

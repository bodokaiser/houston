package model

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DDSParamTestSuite struct {
	suite.Suite
}

func (s *DDSParamTestSuite) TestValidateDDSConst() {
	p1 := DDSConst{Value: -1.0}
	p2 := DDSConst{Value: +1.0}

	assert.Error(s.T(), validate.Struct(p1))
	assert.NoError(s.T(), validate.Struct(p2))
}

func (s *DDSParamTestSuite) TestValidateDDSSweep() {
	p1 := DDSSweep{
		Limits:   [2]float64{0, 100},
		Duration: time.Second,
	}
	p2 := DDSSweep{
		Limits:   [2]float64{0, 100},
		Duration: 0,
	}
	p3 := DDSSweep{
		Limits:   [2]float64{100, 0},
		Duration: time.Second,
	}
	p4 := DDSSweep{
		Limits:   [2]float64{-100, 0},
		Duration: time.Second,
	}

	assert.NoError(s.T(), validate.Struct(p1))
	assert.Error(s.T(), validate.Struct(p2))
	assert.Error(s.T(), validate.Struct(p3))
	assert.Error(s.T(), validate.Struct(p4))
}

func (s *DDSParamTestSuite) TestValidateDDSPlayback() {
	p1 := DDSPlayback{
		Interval: time.Nanosecond,
		Data:     []float64{1, 2, 3},
	}
	p2 := DDSPlayback{
		Interval: 0,
		Data:     []float64{1, 2, 3},
	}
	p3 := DDSPlayback{
		Interval: time.Nanosecond,
	}

	assert.NoError(s.T(), validate.Struct(p1))
	assert.Error(s.T(), validate.Struct(p2))
	assert.Error(s.T(), validate.Struct(p3))
}

func (s *DDSParamTestSuite) TestParamMarshalJSON() {
	p1 := DDSParam{Mode: ModeConst}
	p2 := DDSParam{Mode: ModeSweep}
	p3 := DDSParam{Mode: ModePlayback}

	json1, err1 := json.Marshal(p1)
	assert.NoError(s.T(), err1)
	assert.Equal(s.T(), `{"mode":"const"}`, string(json1))

	json2, err2 := json.Marshal(p2)
	assert.NoError(s.T(), err2)
	assert.Equal(s.T(), `{"mode":"sweep"}`, string(json2))

	json3, err3 := json.Marshal(p3)
	assert.NoError(s.T(), err3)
	assert.Equal(s.T(), `{"mode":"playback"}`, string(json3))
}

func (s *DDSParamTestSuite) TestParamUnmarshalJSON() {
	p1 := DDSParam{}
	p2 := DDSParam{}
	p3 := DDSParam{}

	assert.NoError(s.T(), json.Unmarshal([]byte(`{"mode":"const"}`), &p1))
	assert.NoError(s.T(), json.Unmarshal([]byte(`{"mode":"SWEEP"}`), &p2))
	assert.NoError(s.T(), json.Unmarshal([]byte(`{"mode":"playback"}`), &p3))

	assert.Equal(s.T(), ModeConst, p1.Mode)
	assert.Equal(s.T(), ModeSweep, p2.Mode)
	assert.Equal(s.T(), ModePlayback, p3.Mode)

	assert.Error(s.T(), json.Unmarshal([]byte(`{"mode":"playba"}`), &p3))
	assert.Error(s.T(), json.Unmarshal([]byte(`mode":"playba"}`), &p3))
}

func TestDDSParamTestSuite(t *testing.T) {
	suite.Run(t, new(DDSParamTestSuite))
}

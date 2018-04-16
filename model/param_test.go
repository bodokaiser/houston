package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DDSParamTestSuite struct {
	suite.Suite
}

func (s *DDSParamTestSuite) TestMarshalJSON() {
}

func (s *DDSParamTestSuite) TestUnmarshalJSON() {
	p1 := &DDSParam{}
	p2 := &DDSParam{}
	p3 := &DDSParam{}

	assert.NoError(s.T(), json.Unmarshal([]byte(`{
    "mode": "const",
    "value": 10000
  }`), p1))
	assert.NoError(s.T(), json.Unmarshal([]byte(`{
    "mode": "sweep",
    "limits": [1, 10]
  }`), p2))
	assert.NoError(s.T(), json.Unmarshal([]byte(`{
    "mode": "playback",
    "data": [0, 1, 2, 3]
  }`), p3))
	assert.Error(s.T(), json.Unmarshal([]byte(`{
    "value": 10000
  }`), &DDSParam{}))
	assert.Error(s.T(), json.Unmarshal([]byte(`{
    "mode": "foo",
    "value": 10000
  }`), &DDSParam{}))

	assert.Equal(s.T(), &DDSParam{
		DDSConst: &DDSConst{Value: 10000},
	}, p1)
	assert.Equal(s.T(), &DDSParam{
		DDSSweep: &DDSSweep{Limits: []float64{1, 10}},
	}, p2)
	assert.Equal(s.T(), &DDSParam{
		DDSPlayback: &DDSPlayback{Data: []byte{0, 1, 2, 3}},
	}, p3)
}

func TestDDSParamTestSuite(t *testing.T) {
	suite.Run(t, new(DDSParamTestSuite))
}

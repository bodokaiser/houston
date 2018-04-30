import React, {
  Component
} from 'react'
import {
  LocalForm,
  Control,
  Fieldset
} from 'react-redux-form'
import {connect} from 'react-redux'
import convert from 'convert-units'

import {
  updateDevice
} from '../actions/device'

import {
  RangeInput,
  CheckboxInput,
  InputGroup,
  SelectGroup,
  SelectGroupInput
} from '../components/form'
import {
  CollapsableCard
} from '../components/card'

const ConstGroup = ({ param, value }) => (
  <Fieldset model=".const">
    <div className="row align-items-center">
      <div className="col">
        <Control model=".value" component={RangeInput} />
      </div>
      <div className="col-auto">
        <div className="w-9">
          <Control.text model=".value" component={InputGroup} />
        </div>
      </div>
    </div>
  </Fieldset>
)

const SweepGroup = ({ param }) => (
  <Fieldset model=".sweep">
    <div className="row gutters-xs">
      <div className="col-4">
        <Control.text model=".start" component={InputGroup} />
      </div>
      <div className="col-4">
        <Control.text model=".stop" component={InputGroup} />
      </div>
      <div className="col-4">
        <Control.text model=".duration" component={InputGroup} />
      </div>
    </div>
    <div className="mt-3">
      <Control.checkbox model=".noDwellLow" component={CheckboxInput} label="Hold Low" />
      <Control.checkbox model=".noDwellHigh" component={CheckboxInput} label="Hold High" />
    </div>
  </Fieldset>
)

const PlaybackGroup = ({ param }) => (
  <Fieldset model=".playback">
    <div className="row gutters-xs">
      <div className="col-7">
        <Control.text model=".data" component={InputGroup} />
      </div>
      <div className="col-5">
        <Control.text model=".interval" component={InputGroup} />
      </div>
    </div>
    <div className="mt-3">
      <Control.checkbox model=".trigger" component={CheckboxInput} label="Trigger" />
      <Control.checkbox model=".duplex" component={CheckboxInput} label="Duplex" />
    </div>
  </Fieldset>
)

const ModeGroup = () => (
  <SelectGroup>
    <SelectGroupInput model=".mode" value="const" icon="minus" />
    <SelectGroupInput model=".mode" value="sweep" icon="trending-up" />
    <SelectGroupInput model=".mode" value="playback" icon="activity" />
  </SelectGroup>
)

const ParamGroup = ({ param }) => {
  if (param.mode == 'const') {
    return (<ConstGroup />)
  }
  if (param.mode == 'sweep') {
    return (<SweepGroup />)
  }
  if (param.mode == 'playback') {
    return (<PlaybackGroup />)
  }
}

function required(value) {
  return value && value.length > 0
}

class Device extends Component {

  constructor(props) {
    super(props)

    this.handleSubmit = this.handleSubmit.bind(this)
    this.handleChange = this.handleChange.bind(this)
  }

  handleSubmit() {
    //this.props.dispatch(submitDevice(this.props.device))
  }

  handleChange(device) {
    this.props.dispatch(updateDevice(device))
  }

  handleUpdate(form) {
    //this.props.dispatch(updateDeviceProp(this.props.device, name, value))
  }

  render() {
    const { device } = this.props

    return (
      <CollapsableCard title={device.name}>
        <div className="card-body">
          <p className="text-muted mb-5">{device.description}</p>
          <LocalForm
            initialState={device}
            validators={{
              '': (foo) => {!!console.log(foo)}
            }}
            onUpdate={this.handleUpdate}
            onSubmit={this.handleSubmit}
            onChange={this.handleChange}>
            <div className="form-row">
              <div className="form-group col-sm-12">
                <label className="form-label">
                  Amplitude
                </label>
                <Fieldset model=".amplitude">
                  <ModeGroup />
                  <ParamGroup param={device.amplitude} />
                </Fieldset>
              </div>
              <div className="form-group col-sm-12">
                <label className="form-label">
                  Frequency
                </label>
                <Fieldset model=".frequency">
                  <ModeGroup />
                  <ParamGroup param={device.frequency} />
                </Fieldset>
              </div>
              <div className="form-group col-sm-12">
                <label className="form-label">
                  Phase Offset
                </label>
                <Fieldset model=".phase">
                  <ModeGroup />
                  <ParamGroup param={device.phase} />
                </Fieldset>
              </div>
            </div>
          </LocalForm>
        </div>
        <div className="card-footer text-right">
          <div className="d-flex">
            <button type="button" className="btn btn-outline-secondary">Reset</button>
            <button type="button" className="btn btn-primary ml-auto">Update</button>
          </div>
        </div>
      </CollapsableCard>
    )
  }

}

export default connect()(Device)

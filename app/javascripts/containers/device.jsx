import React, {
  Component,
  Fragment
} from 'react'
import {
  LocalForm,
  Control,
  Fieldset
} from 'react-redux-form'
import {connect} from 'react-redux'

import {DefaultForm} from '../components/form'
import {
  RangeInput,
  CheckboxInput,
  InputGroup,
  SelectGroup
} from '../components/input'
import {SubmitButton} from '../components/button'
import {Card} from '../components/device'

import {
  submitDevice,
  updateDevice
} from '../actions/device'

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
      <div className="col-8">
        <Control.text model=".data" component={InputGroup} />
      </div>
      <div className="col-4">
        <Control.text model=".interval" component={InputGroup} />
      </div>
    </div>
    <div className="mt-3">
      <Control.checkbox model=".trigger" component={CheckboxInput} label="Trigger" />
      <Control.checkbox model=".duplex" component={CheckboxInput} label="Duplex" />
    </div>
  </Fieldset>
)

const ParamGroup = ({ param }) => {
  console.log('param group', param)
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


class Device extends Component {

  constructor(props) {
    super(props)

    this.handleSubmit = this.handleSubmit.bind(this)
    this.handleChange = this.handleChange.bind(this)
  }

  handleSubmit() {
    console.log('handle submit', values)
    //this.props.dispatch(submitDevice(this.props.device))
  }

  handleChange(device) {
    console.log('handle change', device)
    this.props.dispatch(updateDevice(device))
  }

  handleUpdate(form) {
    console.log('handle update', form)
    //this.props.dispatch(updateDeviceProp(this.props.device, name, value))
  }

  render() {
    const { device } = this.props

    return (
      <Card title={device.name}>
        <p className="text-muted mb-5">{device.description}</p>
        <LocalForm
          initialState={device}
          onUpdate={this.handleUpdate}
          onSubmit={this.handleSubmit}
          onChange={this.handleChange}>
          <div className="form-row">
            <div className="form-group col-sm-12">
              <label className="form-label">
                Amplitude
              </label>
              <Fieldset model=".amplitude">
                <SelectGroup model=".mode" />
                <ParamGroup param={device.amplitude} />
              </Fieldset>
            </div>
            <div className="form-group col-sm-12">
              <label className="form-label">
                Frequency
              </label>
              <Fieldset model=".frequency">
                <SelectGroup model=".mode" />
                <ParamGroup param={device.frequency} />
              </Fieldset>
            </div>
            <div className="form-group col-sm-12">
              <label className="form-label">
                Phase Offset
              </label>
              <Fieldset model=".phase">
                <SelectGroup model=".mode" />
                <ParamGroup param={device.phase} />
              </Fieldset>
            </div>
          </div>
        </LocalForm>
      </Card>
    )
  }

}

export default connect()(Device)

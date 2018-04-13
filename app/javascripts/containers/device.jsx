import React, {
  Component,
  Fragment
} from 'react'
import {connect} from 'react-redux'

import {DefaultForm} from '../components/form'
import {InputGroup} from '../components/input'
import {SubmitButton} from '../components/button'

import {
  submitDevice,
  updateDevice
} from '../actions/device'

const DeviceForm = ({ amplitude, frequency, phase, onSubmit, onChange }) => (
  <DefaultForm onSubmit={onSubmit} onChange={onChange}>
    <div className="form-row">
      <div className="form-group col-sm-12">
        <InputGroup name="amplitude" label="Amplitude" type="number"
          min="0" max="100" append="%" value={amplitude} />
      </div>
      <div className="form-group col-sm-12">
        <InputGroup name="frequency" label="Frequency" type="number"
          min="1" max="500"
          append="MHz" value={frequency} />
      </div>
      <div className="form-group col-sm-12">
        <InputGroup name="phase" label="Phase" type="number"
          min="0" max={2*Math.PI}
          append="rad" value={phase} />
      </div>
    </div>
    <div className="form-row">
      <div className="form-group col-sm-12">
        <SubmitButton>Update</SubmitButton>
      </div>
    </div>
  </DefaultForm>
)

class Device extends Component {

  constructor(props) {
    super(props)

    this.handleSubmit = this.handleSubmit.bind(this)
    this.handleChange = this.handleChange.bind(this)
  }

  handleSubmit() {
    this.props.dispatch(submitDevice(this.props.device))
  }

  handleChange(name, value) {
    this.props.device[name] = value
    this.props.dispatch(updateDevice(this.props.device))
  }

  render() {
    const { device } = this.props

    return (
      <div className="card">
        <div className="card-header">
          {device.name}
        </div>
        <div className="card-body">
          <DeviceForm onSubmit={this.handleSubmit}
                      onChange={this.handleChange}
                      {...device} />
        </div>
      </div>
    )
  }

}

export default connect()(Device)

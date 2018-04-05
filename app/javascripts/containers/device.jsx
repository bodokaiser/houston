import React, {
  Component,
  Fragment
} from 'react'
import {connect} from 'react-redux'

import {DefaultForm} from '../components/form'
import {InputGroup} from '../components/input'
import {SubmitButton} from '../components/button'
import {NavTabs} from '../components/nav'
import {
  EditIcon,
  MicrochipIcon
} from '../components/icon'

import {
  updateDeviceMode,
  updateDeviceName
} from '../actions/local'
import {
  fetchDevicesIfNeeded
} from '../actions/remote'

const SingleToneForm = ({ amplitude, frequency }) => (
  <DefaultForm>
    <div className="form-row">
      <div className="form-group col-sm-12">
        <InputGroup name="amplitude" label="Amplitude" type="number"
          append="dB" value={amplitude} />
      </div>
      <div className="form-group col-sm-12">
        <InputGroup name="frequency" label="Frequency" type="number"
          append="Hz" value={frequency} />
      </div>
    </div>
    <div className="form-row">
      <div className="form-group col-sm-12">
        <SubmitButton>Update</SubmitButton>
      </div>
    </div>
  </DefaultForm>
)

const LinearSweepForm = ({ startFrequency, stopFrequency, timerInterval }) => (
  <DefaultForm>
    <div className="form-row">
      <div className="form-group col-sm-12">
        <InputGroup name="startFrequency" type="number" label="Start Frequency"
          append="Hz" value={startFrequency} />
      </div>
      <div className="form-group col-sm-12">
        <InputGroup name="stopFrequency" type="number" label="Stop Frequency"
          append="Hz" value={stopFrequency} />
      </div>
    </div>
    <div className="form-row">
      <div className="form-group col-sm-12">
        <InputGroup name="timerInterval" type="number" label="Timer Interval"
          append="s" value={timerInterval} />
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

    this.handleTabClick = this.handleTabClick.bind(this)
  }

  handleTabClick(mode) {
    this.props.dispatch(updateDeviceMode(this.props.device, mode))
  }

  componentDidMount() {
    const { dispatch } = this.props

    dispatch(fetchDevicesIfNeeded())
  }

  render() {
    const { device, links } = this.props

    links.forEach(link => {
      if (link.name == device.mode) link.active = true
    })

    return (
      <div className="card">
        <div className="card-header">
          <div className="btn-toolbar justify-content-between">
            <InputGroup value={device.name} readOnly={true} />
            <div className="btn-group">
              <button className="btn btn-light" type="button">
                <EditIcon />
              </button>
            </div>
          </div>
        </div>
        <div className="card-body">
          <NavTabs links={links} onClick={this.handleTabClick} />
          { device.mode == 'Single Tone' &&
            <SingleToneForm {...device.params.singleTone} /> }
          { device.mode == 'Linear Sweep' &&
            <LinearSweepForm {...device.params.linearSweep} /> }
        </div>
      </div>
    )
  }

}

const mapState = (state) => {
  const { params } = state

  return {
    links: params.modes.map(mode => ({ name: mode }))
  }
}

export default connect(
  mapState
)(Device)

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

    this.state = { nameEditable: false }
    this.handleTabClick = this.handleTabClick.bind(this)
    this.handleEditClick = this.handleEditClick.bind(this)
    this.handleNameSubmit = this.handleNameSubmit.bind(this)
    this.handleNameChange = this.handleNameChange.bind(this)
  }

  handleEditClick(element) {
    element.preventDefault()

    this.setState({ nameEditable: true })
  }

  handleTabClick(mode) {
    this.props.dispatch(updateDeviceMode(this.props.device, mode))
  }

  handleNameSubmit(element) {
    element.preventDefault()

    this.setState({ nameEditable: false })
  }

  handleNameChange(name) {
    this.props.dispatch(updateDeviceName(this.props.device, name))
  }

  componentDidMount() {
    const { dispatch } = this.props

    dispatch(fetchDevicesIfNeeded())
  }

  render() {
    const { device, links } = this.props
    const { nameEditable } = this.state

    links.forEach(link => {
      if (link.name == device.mode) link.active = true
    })

    return (
      <div className="card">
        <div className="card-header">
          <div className="btn-toolbar justify-content-between">
            <form onSubmit={this.handleNameSubmit}>
              <InputGroup value={device.name} readOnly={!nameEditable} onChange={this.handleNameChange} />
            </form>
            <div className="btn-group" hidden={nameEditable}>
              <button className="btn btn-light" type="button"
                onClick={this.handleEditClick}>
                <EditIcon />
              </button>
            </div>
          </div>
        </div>
        <div className="card-body">
          <NavTabs links={links} onClick={this.handleTabClick} />
          { device.mode == 'Single Tone' &&
            <SingleToneForm {...device.params.singleTone} /> }
          { device.mode == 'Sweep' &&
            <LinearSweepForm {...device.params.sweep} /> }
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

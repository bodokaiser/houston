import React, {
  Component,
  Fragment
} from 'react'
import {connect} from 'react-redux'
import {bindActionCreators} from 'redux'

import {DefaultForm} from '../components/form'
import {
  InputGroup,
  SelectGroup
} from '../components/input'
import {SubmitButton} from '../components/button'
import {NavTabs} from '../components/nav'
import {
  EditIcon,
  MicrochipIcon
} from '../components/icon'

import {
  submitDevice,
  updateDeviceMode,
  updateDeviceName,
  updateDeviceSweep,
  updateDeviceSingleTone
} from '../actions/device'

const SingleToneForm = ({ amplitude, frequency, onChange, onSubmit }) => (
  <form onSubmit={onSubmit}>
    <div className="form-row">
      <div className="form-group col-sm-12">
        <InputGroup name="amplitude" label="Amplitude" type="number"
          append="dB" value={amplitude} onChange={onChange}/>
      </div>
      <div className="form-group col-sm-12">
        <InputGroup name="frequency" label="Frequency" type="number"
          append="Hz" value={frequency} onChange={onChange} />
      </div>
    </div>
    <div className="form-row">
      <div className="form-group col-sm-12">
        <SubmitButton>Update</SubmitButton>
      </div>
    </div>
  </form>
)

const SweepForm = ({ startFrequency, stopFrequency, interval, waveform,
  waveforms, onChange, onSubmit }) => (
  <form onSubmit={onSubmit}>
    <div className="form-row">
      <div className="form-group col-sm-12">
        <InputGroup name="startFrequency" type="number" label="Start Frequency"
          append="Hz" value={startFrequency} onChange={onChange} />
      </div>
      <div className="form-group col-sm-12">
        <InputGroup name="stopFrequency" type="number" label="Stop Frequency"
          append="Hz" value={stopFrequency} onChange={onChange} />
      </div>
    </div>
    <div className="form-row">
      <div className="form-group col-sm-12">
        <InputGroup name="timerInterval" type="number" label="Timer Interval"
          append="s" value={interval} onChange={onChange} />
      </div>
    </div>
    <div className="form-row">
      <div className="form-group col-sm-12">
        <SelectGroup name="waveform" label="Waveform"
          value={waveform} options={waveforms} onChange={onChange} />
      </div>
    </div>
    <div className="form-row">
      <div className="form-group col-sm-12">
        <SubmitButton>Update</SubmitButton>
      </div>
    </div>
  </form>
)

class Device extends Component {

  constructor(props) {
    super(props)

    this.state = { nameEditable: false }
    this.handleTabClick = this.handleTabClick.bind(this)
    this.handleEditClick = this.handleEditClick.bind(this)
    this.handleNameSubmit = this.handleNameSubmit.bind(this)
    this.handleNameChange = this.handleNameChange.bind(this)
    this.handleSingleToneChange = this.handleSingleToneChange.bind(this)
    this.handleSingleToneSubmit = this.handleSingleToneSubmit.bind(this)
    this.handleSweepChange = this.handleSweepChange.bind(this)
  }

  handleEditClick(element) {
    element.preventDefault()

    this.setState({ nameEditable: true })
  }

  handleTabClick(mode) {
    this.props.updateMode(this.props.device, mode)
  }

  handleNameSubmit(element) {
    element.preventDefault()

    this.setState({ nameEditable: false })
  }

  handleNameChange(name) {
    this.props.updateName(this.props.device, name)
  }

  handleSingleToneChange(name, value) {
    var data = {}
    data[name] = value

    this.props.updateSingleTone(this.props.device, data)
  }

  handleSingleToneSubmit(e) {
    e.preventDefault()

    this.props.submitSingleTone(this.props.device)
  }

  handleSweepChange(name, value) {
    this.props.updateSweep(this.props.device, { name: value })
  }

  render() {
    const { device, links, waveforms } = this.props
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
          <div className={(device.mode == 'Single Tone') ? '' : 'd-none'}>
            <SingleToneForm
              onChange={this.handleSingleToneChange}
              onSubmit={this.handleSingleToneSubmit}
              {...device.singleTone} />
          </div>
          <div className={(device.mode == 'Sweep') ? '' : 'd-none'}>
            <SweepForm onChange={this.handleSweepChange} {...{...device.sweep, waveforms}} />
          </div>
        </div>
      </div>
    )
  }

}

const mapState = state => ({
  links: state.specs.AD9910.modes.map(mode => ({ name: mode })),
  waveforms: state.specs.AD9910.modes
})

const mapDispatch = dispatch => ({
  updateName: bindActionCreators(updateDeviceName, dispatch),
  updateMode: bindActionCreators(updateDeviceMode, dispatch),
  updateSweep: bindActionCreators(updateDeviceSweep, dispatch),
  updateSingleTone: bindActionCreators(updateDeviceSingleTone, dispatch),
  submitSingleTone: bindActionCreators(submitDevice, dispatch)
})

export default connect(
  mapState,
  mapDispatch
)(Device)

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
        <label className="form-label">
          Amplitude
        </label>
        <div className="selectgroup w-100">
          <label className="selectgroup-item">
            <input type="radio" name="amplitude" value="const" className="selectgroup-input" />
            <span className="selectgroup-button selectgroup-button-icon">
              <i className="fe fe-minus"></i>
            </span>
          </label>
          <label className="selectgroup-item">
            <input type="radio" name="amplitude" value="sweep" className="selectgroup-input" />
            <span className="selectgroup-button selectgroup-button-icon">
              <i className="fe fe-trending-up"></i>
            </span>
          </label>
          <label className="selectgroup-item">
            <input type="radio" name="amplitude" value="playback" className="selectgroup-input" checked />
            <span className="selectgroup-button selectgroup-button-icon">
              <i className="fe fe-activity"></i>
            </span>
          </label>
        </div>
        <div className="row gutters-xs">
          <div className="col-8">
            <input type="text" className="form-control" placeholder="Data" />
          </div>
          <div className="col-4">
            <input type="text" className="form-control" placeholder="Time" />
          </div>
        </div>
        <div className="mt-3">
          <label className="custom-control custom-checkbox custom-control-inline">
            <input type="checkbox" className="custom-control-input" name="example-inline-checkbox1" value="option1" checked="" />
            <span className="custom-control-label">Trigger</span>
          </label>
          <label className="custom-control custom-checkbox custom-control-inline">
            <input type="checkbox" className="custom-control-input" name="example-inline-checkbox2" value="option2" />
            <span className="custom-control-label">Duplex</span>
          </label>
        </div>
      </div>
      <div className="form-group col-sm-12">
        <label className="form-label">
          Frequency
        </label>
        <div className="selectgroup w-100">
          <label className="selectgroup-item">
            <input type="radio" name="frequency" value="const" className="selectgroup-input" />
            <span className="selectgroup-button selectgroup-button-icon">
              <i className="fe fe-minus"></i>
            </span>
          </label>
          <label className="selectgroup-item">
            <input type="radio" name="frequency" value="sweep" className="selectgroup-input" checked />
            <span className="selectgroup-button selectgroup-button-icon">
              <i className="fe fe-trending-up"></i>
            </span>
          </label>
          <label className="selectgroup-item">
            <input type="radio" name="frequency" value="playback" className="selectgroup-input" />
            <span className="selectgroup-button selectgroup-button-icon">
              <i className="fe fe-activity"></i>
            </span>
          </label>
        </div>
        <div className="row gutters-xs">
          <div className="col-4">
            <input type="text" className="form-control" placeholder="Start" />
          </div>
          <div className="col-4">
            <input type="text" className="form-control" placeholder="Stop" />
          </div>
          <div className="col-4">
            <input type="text" className="form-control" placeholder="Time" />
          </div>
        </div>
        <div className="mt-3">
          <label className="custom-control custom-checkbox custom-control-inline">
            <input type="checkbox" className="custom-control-input" name="example-inline-checkbox1" value="option1" checked />
            <span className="custom-control-label">Hold Low</span>
          </label>
          <label className="custom-control custom-checkbox custom-control-inline">
            <input type="checkbox" className="custom-control-input" name="example-inline-checkbox2" value="option2" checked />
            <span className="custom-control-label">Hold High</span>
          </label>
        </div>
      </div>
      <div className="form-group col-sm-12">
        <label className="form-label">
          Phase Offset
        </label>
        <div className="selectgroup w-100">
          <label className="selectgroup-item">
            <input type="radio" name="phase" value="const" className="selectgroup-input" checked />
            <span className="selectgroup-button selectgroup-button-icon">
              <i className="fe fe-minus"></i>
            </span>
          </label>
          <label className="selectgroup-item">
            <input type="radio" name="phase" value="sweep" className="selectgroup-input" />
            <span className="selectgroup-button selectgroup-button-icon">
              <i className="fe fe-trending-up"></i>
            </span>
          </label>
          <label className="selectgroup-item">
            <input type="radio" name="phase" value="playback" className="selectgroup-input" />
            <span className="selectgroup-button selectgroup-button-icon">
              <i className="fe fe-activity"></i>
            </span>
          </label>
        </div>
        <div className="row align-items-center">
          <div className="col">
            <input type="range" className="form-control custom-range" step="1" min="0" max={2*Math.PI} />
          </div>
          <div className="col-auto">
            <div className="w-9">
              <input type="text" className="form-control" placeholder="0.0 rad" />
            </div>
          </div>
        </div>
      </div>
    </div>
    <div className="card-footer text-right">
      <div className="d-flex pt-2">
        <button type="button" className="btn btn-outline-secondary">Reset</button>
        <button type="button" className="btn btn-primary ml-auto">Update</button>
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
        <div className="card-status bg-blue"></div>
        <div className="card-header">
          <h3 className="card-title">{device.name}</h3>
          <div className="card-options">
            <span className="card-options-collapse">
              <i className="fe fe-edit-2 mr-2"></i>

              <i className="fe fe-chevron-up"></i>
            </span>
          </div>
        </div>
        <div className="card-body">
          <p className="text-muted mb-5">Direct Digital Synthesizer #1</p>
          <DeviceForm onSubmit={this.handleSubmit}
                      onChange={this.handleChange}
                      {...device} />
        </div>
      </div>
    )
  }

}

export default connect()(Device)

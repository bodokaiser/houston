import React, {
  Component,
  Fragment
} from 'react'
import {connect} from 'react-redux'

import {Page} from '../components/layout'
import {DefaultForm} from '../components/form'
import {
  Range,
  Checkbox,
  InputGroup,
  SelectGroup
} from '../components/input'
import {SubmitButton} from '../components/button'
import {Card} from '../components/device'

import {
  submitDevice,
  updateDevice
} from '../actions/device'

const ConstGroup = ({ value }) => (
  <div className="row align-items-center">
    <div className="col">
      <Range step={0.001} min={0} max={2*Math.PI} />
    </div>
    <div className="col-auto">
      <div className="w-9">
        <InputGroup type="text" name="value" value={value} />
      </div>
    </div>
  </div>
)

const SweepGroup = ({ start, stop, duration, nodwells }) => (
  <Fragment>
    <div className="row gutters-xs">
      <div className="col-4">
        <InputGroup type="text" name="start" placeholder="Start" value={start} />
      </div>
      <div className="col-4">
        <InputGroup type="text" name="stop" placeholder="Stop" value={stop} />
      </div>
      <div className="col-4">
        <InputGroup type="text" name="duration" placeholder="Time" value={duration} />
      </div>
    </div>
    <div className="mt-3">
      <Checkbox name="nodwells0" label="Hold Low" checked={nodwells[0]} />
      <Checkbox name="nodwells1" label="Hold High" checked={nodwells[1]} />
    </div>
  </Fragment>
)

const PlaybackGroup = ({ data, interval, trigger, duplex}) => (
  <Fragment>
    <div className="row gutters-xs">
      <div className="col-8">
        <InputGroup type="text" placeholder="Data" value={data} />
      </div>
      <div className="col-4">
        <InputGroup type="text" placeholder="Time" value={interval} />
      </div>
    </div>
    <div className="mt-3">
      <Checkbox name="trigger" label="Trigger" checked={trigger} />
      <Checkbox name="Duplex" label="Duplex" checked={duplex} />
    </div>
  </Fragment>
)

const ParamGroup = ({ param }) => {
  if (param.mode == 'const') {
    return (<ConstGroup {...param.const}/>)
  }
  if (param.mode == 'sweep') {
    return (<SweepGroup {...param.sweep} />)
  }
  if (param.mode == 'playback') {
    return (<PlaybackGroup {...param.playback} />)
  }
}


const DeviceForm = ({ device, onSubmit, onChange }) => (
  <DefaultForm onSubmit={onSubmit} onChange={onChange}>
    <div className="form-row">
      <div className="form-group col-sm-12">
        <label className="form-label">
          Amplitude
        </label>
        <SelectGroup name="amplitude" value={device.amplitude.mode} options={[
          { value: 'const', icon: 'minus' },
          { value: 'sweep', icon: 'trending-up' },
          { value: 'playback', icon: 'activity' }
        ]} />
      </div>
      <ParamGroup param={device.amplitude} />
      <div className="form-group col-sm-12">
        <label className="form-label">
          Frequency
        </label>
        <SelectGroup name="frequency" value={device.frequency.mode} options={[
          { value: 'const', icon: 'minus' },
          { value: 'sweep', icon: 'trending-up' },
          { value: 'playback', icon: 'activity' }
        ]} />
        <ParamGroup param={device.frequency} />
      </div>
      <div className="form-group col-sm-12">
        <label className="form-label">
          Phase Offset
        </label>
        <SelectGroup name="phase" value={device.phase.mode} options={[
          { value: 'const', icon: 'minus' },
          { value: 'sweep', icon: 'trending-up' },
          { value: 'playback', icon: 'activity' }
        ]} />
        <ParamGroup param={device.phase} />
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
      <Card title={device.name}>
        <p className="text-muted mb-5">{device.description}</p>
        <DeviceForm onSubmit={this.handleSubmit}
                    onChange={this.handleChange}
                    device={device} />
      </Card>
    )
  }

}

class Devices extends Component {

  render() {
    const { devices } = this.props

    return (
      <Page title="Devices">
        <div className="row">
        {devices.map((device, index) => (
          <div className="col-8 offset-2 col-sm-6 offset-sm-0 col-md-5 col-lg-4 col-xl-3" key={index}>
            <Device key={index} device={device} />
          </div>
        ))}
        </div>
      </Page>
    )
  }

}

const mapState = state => ({
  devices: state.devices
})

export default connect(
  mapState
)(Devices)

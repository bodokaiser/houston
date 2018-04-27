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

const ConstGroup = () => (
  <div className="row align-items-center">
    <div className="col">
      <Range step={0.001} min={0} max={2*Math.PI} />
    </div>
    <div className="col-auto">
      <div className="w-9">
        <InputGroup type="text" name="phase" placeholder="0.0 rad" />
      </div>
    </div>
  </div>
)

const SweepGroup = ({}) => (
  <Fragment>
    <div className="row gutters-xs">
      <div className="col-4">
        <InputGroup type="text" placeholder="Start" />
      </div>
      <div className="col-4">
        <InputGroup type="text" placeholder="Stop" />
      </div>
      <div className="col-4">
        <InputGroup type="text" placeholder="Time" />
      </div>
    </div>
    <div className="mt-3">
      <Checkbox name="nodwellLow" label="Hold Low" value={false} checked={false} />
      <Checkbox name="nodwellHigh" label="Hold High" value={false} checked={false} />
    </div>
  </Fragment>
)

const PlaybackGroup = ({ }) => (
  <Fragment>
    <div className="row gutters-xs">
      <div className="col-8">
        <InputGroup type="text" placeholder="Data" />
      </div>
      <div className="col-4">
        <InputGroup type="text" placeholder="Data" />
      </div>
    </div>
    <div className="mt-3">
      <Checkbox name="trigger" label="Trigger" value={true} checked={false} />
      <Checkbox name="Duplex" label="Duplex" value={true} checked={true} />
    </div>
  </Fragment>
)


const DeviceForm = ({ amplitude, frequency, phase, onSubmit, onChange }) => (
  <DefaultForm onSubmit={onSubmit} onChange={onChange}>
    <div className="form-row">
      <div className="form-group col-sm-12">
        <label className="form-label">
          Amplitude
        </label>
        <SelectGroup name="amplitude" value="playback" options={[
          { value: 'const', icon: 'minus' },
          { value: 'sweep', icon: 'trending-up' },
          { value: 'playback', icon: 'activity' }
        ]} />
      </div>
      <PlaybackGroup />
      <div className="form-group col-sm-12">
        <label className="form-label">
          Frequency
        </label>
        <SelectGroup name="frequency" value="sweep" options={[
          { value: 'const', icon: 'minus' },
          { value: 'sweep', icon: 'trending-up' },
          { value: 'playback', icon: 'activity' }
        ]} />
        <SweepGroup />
      </div>
      <div className="form-group col-sm-12">
        <label className="form-label">
          Phase Offset
        </label>
        <SelectGroup name="phase" value="const" options={[
          { value: 'const', icon: 'minus' },
          { value: 'sweep', icon: 'trending-up' },
          { value: 'playback', icon: 'activity' }
        ]} />
        <ConstGroup />
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
        <p className="text-muted mb-5">Direct Digital Synthesizer #1</p>
        <DeviceForm onSubmit={this.handleSubmit}
                    onChange={this.handleChange}
                    {...device} />
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

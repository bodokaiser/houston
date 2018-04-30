import React, {
  Component
} from 'react'
import {
  LocalForm,
  Fieldset
} from 'react-redux-form'
import {connect} from 'react-redux'

import {
  updateDevice
} from '../actions/device'

import {
  modes,
  required,
  quantity,
  quantities
} from '../validators/device'

import {
  InputGroup,
  SelectGroup,
  SelectGroupOption,
  CheckboxGroup
} from '../components/form/group'
import {
  TextInput,
  RangeInput
} from '../components/form/input'
import {
  CollapsableCard
} from '../components/card'

const ConstGroup = ({ measure }) => (
  <Fieldset model=".const">
    <div className="row align-items-center">
      <div className="col">
        <InputGroup
          model=".value"
          validators={{
            required: required(),
            quantity: quantity(measure)
          }} />
      </div>
    </div>
  </Fieldset>
)

const SweepGroup = ({ measure }) => (
  <Fieldset model=".sweep">
    <div className="row gutters-xs">
      <div className="col-4">
        <InputGroup
          model=".start"
          validators={{
            required: required(),
            quantity: quantity(measure)
          }} />
      </div>
      <div className="col-4">
        <InputGroup
          model=".stop"
          validators={{
            required: required(),
            quantity: quantity(measure)
          }} />
      </div>
      <div className="col-4">
        <InputGroup
          model=".duration"
          validators={{
            required: required(),
            quantity: quantity('time')
          }} />
      </div>
    </div>
    <div className="mt-3">
      <CheckboxGroup model=".noDwellLow" label="Hold Low" />
      <CheckboxGroup model=".noDwellHigh" label="Hold High" />
    </div>
  </Fieldset>
)

const PlaybackGroup = ({ measure }) => (
  <Fieldset model=".playback">
    <div className="row gutters-xs">
      <div className="col-7">
        <InputGroup
          model=".data"
          validators={{
            required: required(),
            quantity: quantities(measure)
          }}
        />
      </div>
      <div className="col-5">
        <InputGroup
          model=".interval"
          validators={{
            required: required(),
            quantity: quantity('time')
          }} />
      </div>
    </div>
    <div className="mt-3">
      <CheckboxGroup model=".noDwellLow" label="Hold Low" />
      <CheckboxGroup model=".noDwellHigh" label="Hold High" />
    </div>
  </Fieldset>
)

const ModeGroup = () => (
  <SelectGroup model="local.mode">
    <SelectGroupOption model=".mode" value="const" icon="minus" />
    <SelectGroupOption model=".mode" value="sweep" icon="trending-up" />
    <SelectGroupOption model=".mode" value="playback" icon="activity" />
  </SelectGroup>
)

class Device extends Component {

  constructor(props) {
    super(props)

    this.state = {}
  }

  handleSubmit() {
    //this.props.dispatch(submitDevice(this.props.device))
  }

  handleChange(device) {
    this.props.dispatch(updateDevice(device))
  }

  handleUpdate(form) {
    this.setState({ form })
  }

  render() {
    const { device } = this.props
    const { form } = this.state

    var alert  = form && form.$form.errors.mode &&
      'You can only use one sweep and one playback at a time.'

    return (
      <CollapsableCard title={device.name} alert={alert}>
        <div className="card-body">
          <p className="text-muted mb-5">{device.description}</p>
          <LocalForm
            initialState={device}
            validators={{
              '': {
                mode: modes()
              }
            }}
            onUpdate={form => this.handleUpdate(form)}
            onSubmit={device => this.handleSubmit(device)}
            onChange={device => this.handleChange(device)}>
            <div className="form-row">
              <div className="form-group col-sm-12">
                <label className="form-label">
                  Amplitude
                </label>
                <Fieldset model=".amplitude">
                  <ModeGroup />
                  { device.amplitude.mode == 'const' &&
                  <ConstGroup measure="relative" /> }
                  { device.amplitude.mode == 'sweep' &&
                  <SweepGroup measure="relative" /> }
                  { device.amplitude.mode == 'playback' &&
                  <PlaybackGroup measure="relative" /> }
                </Fieldset>
              </div>
              <div className="form-group col-sm-12">
                <label className="form-label">
                  Frequency
                </label>
                <Fieldset model=".frequency">
                  <ModeGroup />
                  { device.frequency.mode == 'const' &&
                  <ConstGroup measure="frequency" /> }
                  { device.frequency.mode == 'sweep' &&
                  <SweepGroup measure="frequency" /> }
                  { device.frequency.mode == 'playback' &&
                  <PlaybackGroup measure="frequency" /> }
                </Fieldset>
              </div>
              <div className="form-group col-sm-12">
                <label className="form-label">
                  Phase Offset
                </label>
                <Fieldset model=".phase">
                  <ModeGroup />
                  { device.phase.mode == 'const' &&
                  <ConstGroup measure="angle" /> }
                  { device.phase.mode == 'sweep' &&
                  <SweepGroup measure="angle" /> }
                  { device.phase.mode == 'playback' &&
                  <PlaybackGroup measure="angle" /> }
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

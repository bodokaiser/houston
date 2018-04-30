import React, {
  Component
} from 'react'
import {
  LocalForm,
  Fieldset
} from 'react-redux-form'
import {connect} from 'react-redux'
import {
  isEmpty
} from 'validator'
import convert from 'convert-units'

import {
  updateDevice
} from '../actions/device'

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

const ConstGroup = ({ param, value }) => (
  <Fieldset model=".const">
    <div className="row align-items-center">
      <div className="col">
        <RangeInput model=".value" />
      </div>
      <div className="col-auto">
        <div className="w-9">
          <InputGroup
            model=".value"
            validators={{
              required: v => !isEmpty(v),

            }} />
        </div>
      </div>
    </div>
  </Fieldset>
)

const SweepGroup = ({ param }) => (
  <Fieldset model=".sweep">
    <div className="row gutters-xs">
      <div className="col-4">
        <InputGroup model=".start" />
      </div>
      <div className="col-4">
        <InputGroup model=".stop" />
      </div>
      <div className="col-4">
        <InputGroup model=".duration" />
      </div>
    </div>
    <div className="mt-3">
      <CheckboxGroup model=".noDwellLow" label="Hold Low" />
      <CheckboxGroup model=".noDwellHigh" label="Hold High" />
    </div>
  </Fieldset>
)

const PlaybackGroup = ({  }) => (
  <Fieldset model=".playback">
    <div className="row gutters-xs">
      <div className="col-7">
        <InputGroup model=".data" />
      </div>
      <div className="col-5">
        <InputGroup model=".interval" />
      </div>
    </div>
    <div className="mt-3">
      <CheckboxGroup model=".noDwellLow" label="Hold Low" />
      <CheckboxGroup model=".noDwellHigh" label="Hold High" />
    </div>
  </Fieldset>
)

const ModeGroup = () => (
  <SelectGroup>
    <SelectGroupOption model=".mode" value="const" icon="minus" />
    <SelectGroupOption model=".mode" value="sweep" icon="trending-up" />
    <SelectGroupOption model=".mode" value="playback" icon="activity" />
  </SelectGroup>
)

class Device extends Component {

  constructor(props) {
    super(props)
  }

  handleSubmit() {
    //this.props.dispatch(submitDevice(this.props.device))
  }

  handleChange(device) {
    this.props.dispatch(updateDevice(device))
  }

  handleUpdate(form) {
  }

  render() {
    const { device } = this.props

    return (
      <CollapsableCard title={device.name}>
        <div className="card-body">
          <p className="text-muted mb-5">{device.description}</p>
          <LocalForm
            initialState={device}
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
                  <ConstGroup  /> }
                  { device.amplitude.mode == 'sweep' &&
                  <SweepGroup /> }
                  { device.amplitude.mode == 'playback' &&
                  <PlaybackGroup /> }
                </Fieldset>
              </div>
              <div className="form-group col-sm-12">
                <label className="form-label">
                  Frequency
                </label>
                <Fieldset model=".frequency">
                  <ModeGroup />
                  { device.frequency.mode == 'const' &&
                  <ConstGroup /> }
                  { device.frequency.mode == 'sweep' &&
                  <SweepGroup /> }
                  { device.frequency.mode == 'playback' &&
                  <PlaybackGroup /> }
                </Fieldset>
              </div>
              <div className="form-group col-sm-12">
                <label className="form-label">
                  Phase Offset
                </label>
                <Fieldset model=".phase">
                  <ModeGroup />
                  { device.phase.mode == 'const' &&
                  <ConstGroup /> }
                  { device.phase.mode == 'sweep' &&
                  <SweepGroup /> }
                  { device.phase.mode == 'playback' &&
                  <PlaybackGroup /> }
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

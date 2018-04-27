import React, {Fragment} from 'react'

import {
  LocalForm,
  Control
} from 'react-redux-form'

export const InputGroup = (props) => (
  <Fragment>
    { props.label &&
    <label className="form-label" htmlFor={props.name}>
      { props.label }
    </label> }
    <div className="input-group">
      { props.prepend &&
      <div className="input-group-prepend">
        <div className="input-group-text">{ props.prepend }</div>
      </div> }
      <input className="form-control" {...props} />
      <div className="valid-feedback">{ props.validation }</div>
      { props.append &&
      <div className="input-group-append">
        <div className="input-group-text">{ props.append }</div>
      </div> }
    </div>
  </Fragment>
)

export const SelectGroupInput = (props) => (
  <input className="selectgroup-input" type="radio" {...props} />
)

export const SelectGroup = ({ model }) => (
  <div className="selectgroup w-100">
    <label className="selectgroup-item">
      <Control.radio model={model} component={SelectGroupInput} value="const" />
      <span className="selectgroup-button selectgroup-button-icon">
        <i className={`fe fe-minus`}></i>
      </span>
    </label>
    <label className="selectgroup-item">
      <Control.radio model={model} component={SelectGroupInput} value="sweep" />
      <span className="selectgroup-button selectgroup-button-icon">
        <i className={`fe fe-trending-up`}></i>
      </span>
    </label>
    <label className="selectgroup-item">
      <Control.radio model={model} component={SelectGroupInput} value="playback" />
      <span className="selectgroup-button selectgroup-button-icon">
        <i className={`fe fe-activity`}></i>
      </span>
    </label>
  </div>
)

export const RangeInput = (props) => (
  <input className="form-control custom-range" type="range" {...props} />
)


export const CheckboxInput = (props) => (
  <label className="custom-control custom-checkbox custom-control-inline">
    <input className="custom-control-input" type="checkbox" {...props} />
    <span className="custom-control-label">{ props.label }</span>
  </label>
)

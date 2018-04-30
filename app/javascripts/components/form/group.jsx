import React, {Fragment} from 'react'

import {
  TextInput,
  SelectInput,
  CheckboxInput
} from './input'
import {
  InvalidFeedback
} from './feedback'

export const InputGroup = ({ model, label, prepend, append, value, validators, messages }) => (
  <Fragment>
    { label &&
    <label className="form-label" htmlFor={model}>
      { label }
    </label> }
    <div className="input-group">
      { prepend &&
      <div className="input-group-prepend">
        <div className="input-group-text">{ prepend }</div>
      </div> }
      <TextInput model={model} validators={validators} />
      <InvalidFeedback model={model} />
      { append &&
      <div className="input-group-append">
        <div className="input-group-text">{ append }</div>
      </div> }
    </div>
  </Fragment>
)

export const SelectGroup = ({ children }) => (
  <Fragment>
    <div className="selectgroup w-100">
      { children }
    </div>
  </Fragment>
)

export const SelectGroupOption = ({ model, icon, value }) => (
  <label className="selectgroup-item">
    <SelectInput model={model} value={value} />
    { icon &&
    <span className="selectgroup-button selectgroup-button-icon">
      <i className={`fe fe-${icon}`}></i>
    </span> }
  </label>
)

export const CheckboxGroup = ({ model, label }) => (
  <label className="custom-control custom-checkbox custom-control-inline">
    <CheckboxInput model={model} />
    <span className="custom-control-label">{ label }</span>
  </label>
)

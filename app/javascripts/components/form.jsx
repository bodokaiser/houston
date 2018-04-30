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

export const SelectInput = (props) => (
  <input className="selectgroup-input" type="radio" {...props} />
)

export const SelectGroup = ({ children }) => (
  <div className="selectgroup w-100">
    { children }
  </div>
)

export const SelectGroupInput = (props) => (
  <label className="selectgroup-item">
    <Control.radio component={SelectInput} {...props} />
    {props.icon &&
    <span className="selectgroup-button selectgroup-button-icon">
      <i className={`fe fe-${props.icon}`}></i>
    </span>}
  </label>
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

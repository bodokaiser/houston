import React, {Fragment} from 'react'

export const InputGroup = ({ name, type, value, label, append, prepend,
  min, max, step, validation, placeholder }) => (
  <Fragment>
    { label &&
    <label className="form-label" htmlFor={name}>
      { label }
    </label> }
    <div className="input-group">
      { prepend &&
      <div className="input-group-prepend">
        <div className="input-group-text">{ prepend }</div>
      </div> }
      <input className="form-control"
        id={name} type={type || 'text'} defaultValue={value}
        min={min} max={max} step={step} placeholder={placeholder} />
      <div className="valid-feedback">{ validation }</div>
      { append &&
      <div className="input-group-append">
        <div className="input-group-text">{ append }</div>
      </div> }
    </div>
  </Fragment>
)

export const SelectGroup = ({ name, value, options, validation, onChange }) => (
  <div className="selectgroup w-100">
    {options.map((option, index) => (
    <label className="selectgroup-item" key={index}>
      <input className="selectgroup-input" type="radio" name={name}
        value={option.value} checked={option.value == value} />
      {option.icon &&
      <span className="selectgroup-button selectgroup-button-icon">
        <i className={`fe fe-${option.icon}`}></i>
      </span>}
    </label>
    ))}
  </div>
)

export const Range = ({ name, step, min, max }) => (
  <input className="form-control custom-range" type="range"
    step="1" min="0" max={2*Math.PI} />
)


export const Checkbox = ({ name, label, value, checked }) => (
  <label className="custom-control custom-checkbox custom-control-inline">
    <input className="custom-control-input" type="checkbox"
      name={name} value={value} checked={checked} />
    <span className="custom-control-label">{ label }</span>
  </label>
)

import React, {Fragment} from 'react'

export const InputGroup = ({ name, type, value, label, append, prepend,
  min, max, step, validation, placeholder }) => (
  <Fragment>
    { label &&
    <label className="form-label" htmlFor={name}>
      { label }
    </label> }
    <div className="btn-group">
      <button type="button" className="btn btn-secondary">Const</button>
      <button type="button" className="btn btn-secondary">Sweep</button>
      <button type="button" className="btn btn-secondary">Playback</button>
    </div>
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

export const SelectGroup = ({ name, label, value, options, validation, onChange }) => (
  <Fragment>
    { label &&
    <label htmlFor={name}>
      { label }
    </label> }
    <div className="input-group">
      <select className="custom-select" defaultValue={value} onChange={e => onChange(e.target.value)}>
        { options.map((option, index) => (
          <option key={index} value={option}>
            {option}
          </option>
        ))}
      </select>
      <div className="valid-feedback">{ validation }</div>
    </div>
  </Fragment>
)

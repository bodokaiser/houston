import React, {Fragment} from 'react'

export const InputGroup = ({ name, type, value, label, append, prepend,
  validation, placeholder, readOnly, onChange }) => (
  <Fragment>
    { label &&
    <label htmlFor={name}>
      { label }
    </label> }
    <div className="input-group">
      { prepend &&
      <div className="input-group-prepend">
        <div className="input-group-text">{ prepend }</div>
      </div> }
      <input className={`form-control${readOnly ? '-plaintext' : ''}`}
        id={name}
        type={type || 'text'}
        defaultValue={value}
        readOnly={readOnly}
        onChange={e => onChange(e.target.name, e.target.value, e.target)}
        placeholder={placeholder} />
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

import React, {Fragment} from 'react'

export const InputGroup = ({ name, type, value, label, append, prepend, validation, placeholder, readOnly }) => (
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
        placeholder={placeholder} />
      <div className="valid-feedback">{ validation }</div>
      { append &&
      <div className="input-group-append">
        <div className="input-group-text">{ append }</div>
      </div> }
    </div>
  </Fragment>
)

import React from 'react'

export const SubmitButton = ({ children }) => (
  <button type="submit" className="btn btn-primary">
    { children }
  </button>
)

export const ClickButton = ({ children, className, onClick }) => (
  <button type="click"
    className={`btn btn-primary ${className || ''}`}
    onClick={e => onClick()}>
    { children }
  </button>
)

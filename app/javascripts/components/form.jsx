import React from 'react'

export const DefaultForm = ({ children, onChange, onSubmit }) => (
  <form onSubmit={mediateSubmit(onSubmit)}
        onChange={mediateChange(onChange)}>
    { children }
  </form>
)

export const InlineForm = ({ children, onChange, onSubmit }) => (
  <form className="form-inline"
        onSubmit={mediateSubmit(onSubmit)}
        onChange={mediateChange(onChange)}>
    { children }
  </form>
)

function mediateSubmit(callback) {
  return event => {
    event.preventDefault()

    callback()
  }
}

function mediateChange(callback) {

  return event => {
    var name = event.target.name
    var value = event.target.value

    if (event.target.type == 'number') value = parseFloat(value)

    callback(name, value)
  }
}

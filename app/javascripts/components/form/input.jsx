import React from 'react'

import {
  Control
} from 'react-redux-form'

const classNameRange = classNameMapper('form-control custom-range')
const classNameSelect = classNameMapper('selectgroup-input')
const classNameCustom = classNameMapper('custom-control-input')
const classNameControl = classNameMapper('form-control')

export const TextInput = (props) => (
  <Control.text mapProps={{ className: classNameControl }} {...props} />
)

export const RangeInput = (props) => (
  <Control type="range" mapProps={{ className: classNameRange }} {...props} />
)

export const SelectInput = (props) => (
  <Control.radio mapProps={{ className: classNameSelect }} {...props} />
)

export const CheckboxInput = (props) => (
  <Control.checkbox mapProps={{ className: classNameCustom }} {...props} />
)

function classNameMapper(str) {
  return function mapClassName(form) {
    if (!form.fieldValue.valid) {
      str += ' is-invalid'
    }

    return str
  }
}

import React from 'react'

import {
  Errors
} from 'react-redux-form'

const wrapper = ({ children }) => (
  <div className="invalid-feedback">
    { children }
  </div>
)

const messages = {
  required: 'input missing',
  quantity: 'invalid quantity'
}

export const InvalidFeedback = ({ model }) => (
  <Errors
    model={model}
    wrapper={wrapper}
    messages={messages} />
)

import React from 'react'
import ReactDOM from 'react-dom'

import Redux from './redux'

const render = Component => {
  const Redux = require('./redux').default

  ReactDOM.render(<Redux />, document.querySelector('main'))
}

render()

if (module.hot) module.hot.accept(render)

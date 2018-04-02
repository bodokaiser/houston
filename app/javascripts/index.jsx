import React from 'react'
import ReactDOM from 'react-dom'

import App from './containers/app'

import {configureStore} from './store'

if (process.env.NODE_ENV == 'developement') {
  var state = {
    devices: [
      {name: 'Signal Generator 1a', mode: 'const'},
      {name: 'Signal Generator 1b', mode: 'const'},
      {name: 'Signal Generatsor 2a', mode: 'sweep'},
      {name: 'Signal Generator 2b', mode: 'const'}
    ]
  }
} else {
  var state = {}
}

const store = configureStore(state)


ReactDOM.render(<App store={store} />, document.querySelector('main'))

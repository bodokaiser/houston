import React from 'react'
import {render} from 'react-dom'

import App from './containers/app'
import {configureStore} from './store'

const store = configureStore()

render(
  <App store={store} />,
  document.querySelector('content'))

import React from 'react'
import {Provider} from 'react-redux'

import store from './store'

import App from './containers/app'

const Redux = () => (
  <Provider store={store}>
    <App />
  </Provider>
)

export default Redux

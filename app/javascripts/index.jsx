import React from 'react'
import ReactDOM from 'react-dom'
import {Provider} from 'react-redux'

import store from './store'

const render = Component => {
  const App = require('./containers/app').default

  ReactDOM.render(
    <Provider store={store}>
      <App />
    </Provider>,
    document.querySelector('main'))
}

render()

if (module.hot) module.hot.accept(render)

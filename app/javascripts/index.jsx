import React from 'react'
import {render} from 'react-dom'
import {Provider} from 'react-redux'
import {
  BrowserRouter as Router,
  Route,
  Switch
} from 'react-router-dom'

import {configureStore} from './store'

import Navbar from './components/navbar'
import Devices from './containers/devices'

if (process.env.NODE_ENV == 'developement') {
  var state = {
    devices: [
      {name: 'Signal Generator 1a', mode: 'const'},
      {name: 'Signal Generator 1b', mode: 'const'},
      {name: 'Signal Generator 2a', mode: 'sweep'}
    ]
  }
} else {
  var state = {}
}

render(
  <Provider store={configureStore(state)}>
    <Router>
      <content>
        <Navbar />
        <Switch>
          <Route exact path="/" render={() => <h1>Index</h1>} />
          <Route path="/devices" component={Devices} />
        </Switch>
      </content>
    </Router>
  </Provider>,
  document.querySelector('main'))

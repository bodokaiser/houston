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

render(
  <Provider store={configureStore()}>
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

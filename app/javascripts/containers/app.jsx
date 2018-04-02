import React from 'react'
import {
  Provider
} from 'react-redux'
import {
  BrowserRouter as Router,
  Route,
  Switch
} from 'react-router-dom'

import Navbar from '../components/navbar'
import Devices from './devices'

const App = ({ store }) => (
  <Provider store={store}>
    <Router>
      <section>
        <Navbar />
        <Switch>
          <Route exact path="/" render={() => <h1>Index</h1>} />
          <Route path="/devices" component={Devices} />
        </Switch>
      </section>
    </Router>
  </Provider>
)

export default App

import React from 'react'
import {Provider} from 'react-redux'
import {
  BrowserRouter as Router,
  Route,
  Switch
} from 'react-router-dom'

import Navbar from '../components/navbar'
import Home from './home'
import Devices from './devices'

const App = ({ store }) => (
  <Provider store={store}>
    <Router>
      <content>
        <Navbar />
        <Switch>
          <Route exact path="/" component={Home} />
          <Route path="/devices" component={Devices} />
        </Switch>
      </content>
    </Router>
  </Provider>
)

export default App

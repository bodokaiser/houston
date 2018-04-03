import React, { Component }  from 'react'
import {
  BrowserRouter as Router,
  Route,
  Switch
} from 'react-router-dom'

import Home from './home'
import Devices from './devices'
import Navbar from './navbar'

class App extends Component {

  render() {
    return (
      <Router>
        <content>
          <Navbar />
          <Switch>
            <Route exact path="/" component={Home} />
            <Route path="/devices" component={Devices} />
          </Switch>
        </content>
      </Router>
    )
  }

}

export default App

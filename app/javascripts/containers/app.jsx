import React, {
  Component,
  Fragment
} from 'react'
import {
  Route,
  Switch,
  withRouter
} from 'react-router-dom'
import {connect} from 'react-redux'

import Devices from './devices'
import Dashboard from './dashboard'

import {
  Header,
  PageNotFound
} from '../components/layout'
import {ClickButton} from '../components/button'

import {fetchDevicesLazy} from '../actions/device'

class App extends Component {

  componentDidMount() {
    //this.props.dispatch(fetchDevicesLazy())
  }

  render() {
    const { device, devices, isUpdating, isFetching } = this.props

    return (
      <Fragment>
        <Header />
        <div className="container mt-5">
          <Switch>
            <Route path="/" component={Dashboard} exact />
            <Route path="/devices" component={Devices} />
            <Route component={PageNotFound}  />
          </Switch>
        </div>
      </Fragment>
    )
  }

}

const mapState = state => ({
  devices: state.devices
})

export default withRouter(connect(
  mapState
)(App))

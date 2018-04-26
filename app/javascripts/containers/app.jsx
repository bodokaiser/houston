import React, {
  Component,
  Fragment
} from 'react'
import {
  Route,
  Link,
  withRouter
} from 'react-router-dom'
import {connect} from 'react-redux'

import Device from './device'
import {Navbar} from '../components/nav'
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
        <header className="header collapse d-lg-flex p-0">
          <div className="container">
            <div className="row align-items-center">
              <div className="col-lg order-lg-first">
                <ul className="nav nav-tabs border-0 flex-column flex-lg-row">
                  <li className="nav-item">
                    <Link className="nav-link" to="/">
                      <i className="fe fe-home"></i> Home
                    </Link>
                  </li>
                  <li className="nav-item">
                    <Link className="nav-link" to="/devices">
                      <i className="fe fe-cpu"></i> Devices
                    </Link>
                  </li>
                  <li className="nav-item">
                    <a className="nav-link" href="https://github.com/bodokaiser/houston">
                      <i className="fe fe-github"></i> Source
                    </a>
                  </li>
                  <li className="nav-item">
                    <a className="nav-link" href="https://godoc.org/github.com/bodokaiser/houston">
                      <i className="fe fe-file-text"></i> Documentation
                    </a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </header>
        <div className="container mt-5">
          <Route exact path="/" render={() => (
            <h1>Hello World</h1>
          )} />
          <Route path="/devices" render={() => (
            <div className="row">
            {devices.map((device, index) => (
              <div className="col-6 col-sm-6 col-md-4 col-lg-3" key={index}>
                <Device key={index} device={device} />
              </div>
            ))}
            </div>
          )} />
        </div>
      </Fragment>
    )
  }

}

const mapState = state => ({
  device: state.device.device,
  devices: state.device.devices,
  isUpdating: state.device.isUpdating,
  isFetching: state.device.isFetching
})

export default withRouter(connect(
  mapState
)(App))

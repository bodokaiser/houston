import React, {
  Component,
  Fragment
} from 'react'
import {connect} from 'react-redux'
import {bindActionCreators} from 'redux'

import Device from './device'
import {Navbar} from '../components/nav'

import {fetchSpecsLazy} from '../actions/spec'
import {fetchDevicesLazy} from '../actions/device'


class App extends Component {

  componentDidMount() {
    this.props.fetchSpecsLazy()
    this.props.fetchDevicesLazy()
  }

  render() {
    const { specs, devices } = this.props

    return (
      <Fragment>
        <Navbar title="Beagle" />
        <div className="container mt-5">
          <div className="row">
            {devices.map((device, index) => (
              <div className="col-sm-3" key={index}>
                <Device key={index} spec={specs.AD9910} device={device} />
              </div>
            ))}
          </div>
        </div>
      </Fragment>
    )
  }

}

const mapState = state => ({
  specs: state.specs,
  devices: state.devices
})

const mapDispatch = dispatch => ({
    fetchSpecsLazy: bindActionCreators(fetchSpecsLazy, dispatch),
    fetchDevicesLazy: bindActionCreators(fetchDevicesLazy, dispatch)
})

export default connect(
  mapState,
  mapDispatch
)(App)

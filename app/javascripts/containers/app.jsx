import React, {
  Component,
  Fragment
} from 'react'
import {connect} from 'react-redux'
import {bindActionCreators} from 'redux'

import Device from './device'
import {Navbar} from '../components/nav'

import {fetchDevicesLazy} from '../actions/device'


class App extends Component {

  componentDidMount() {
    this.props.fetchDevicesLazy()
  }

  render() {
    const { devices } = this.props

    return (
      <Fragment>
        <Navbar title="Beagle" />
        <div className="container mt-5">
          <div className="row">
            {devices.map((device, index) => (
              <div className="col-sm-3" key={index}>
                <Device key={index} device={device} />
              </div>
            ))}
          </div>
        </div>
      </Fragment>
    )
  }

}

const mapState = state => ({
  devices: state.devices
})

const mapDispatch = dispatch => ({
    fetchDevicesLazy: bindActionCreators(fetchDevicesLazy, dispatch)
})

export default connect(
  mapState,
  mapDispatch
)(App)

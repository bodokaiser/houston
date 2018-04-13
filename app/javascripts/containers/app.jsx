import React, {
  Component,
  Fragment
} from 'react'
import {connect} from 'react-redux'
import {bindActionCreators} from 'redux'

import Device from './device'
import {Navbar} from '../components/nav'
import {ClickButton} from '../components/button'

import {fetchDevicesLazy} from '../actions/device'


class App extends Component {

  componentDidMount() {
    this.props.fetchDevicesLazy()
  }

  render() {
    const { devices } = this.props

    return (
      <Fragment>
        <Navbar title="Houston" />
        <div className="container mt-5">
          <div className="row justify-content-center">
            <div className="col-sm-8 align-self-center">
              { !devices.isFetching && devices.length == 0 &&
                <div class="alert alert-danger">
                  <h4 class="alert-heading">Failed to receive devices.</h4>
                  <p>
                    Check if the device server is running and if the app uses the
                    correct url.
                  </p>
                  <ClickButton className="btn-danger"
                    onClick={this.props.fetchDevicesLazy}>
                    Retry
                  </ClickButton>
                </div> }
              { devices.isFetching &&
                <p className="lead mt-5 text-center">
                  Fetching available devices for you ...
                </p> }
            </div>
          </div>
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

import React, {
  Component,
  Fragment
} from 'react'
import {connect} from 'react-redux'

import {Page} from '../components/layout'
import Device from './device'

class Devices extends Component {

  render() {
    const { devices } = this.props

    return (
      <Page title="Devices">
        <div className="row">
        {devices.map((device, index) => (
          <div className="col-8 offset-2 col-sm-6 offset-sm-0 col-md-5 col-lg-4 col-xl-3" key={index}>
            <Device key={index} device={device} />
          </div>
        ))}
        </div>
      </Page>
    )
  }

}

const mapState = state => ({
  devices: state.devices
})

export default connect(
  mapState
)(Devices)

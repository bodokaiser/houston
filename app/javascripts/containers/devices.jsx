import React, { Component } from 'react'
import { connect } from 'react-redux'

import {
  fetchDevicesIfNeeded,
  requestDevices,
  receiveDevices
} from '../actions/device'

class Devices extends Component {

  render() {
    return (
      <section>
        <h2>Devices</h2>
      </section>
    )
  }

  componentDidMount() {
    const { dispatch } = this.props

    dispatch(fetchDevicesIfNeeded())
  }

}

const mapState = (state) => (
  { devices: state.devices }
)

export default connect(
  mapState
)(Devices)

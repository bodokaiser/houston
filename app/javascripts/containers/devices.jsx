import React, { Component } from 'react'
import { connect } from 'react-redux'

import Layout from '../components/layout'

import {
  fetchDevicesIfNeeded,
  requestDevices,
  receiveDevices
} from '../actions/device'


class Devices extends Component {

  render() {
    const { devices } = this.props

    return (
      <Layout className="col-sm">
        <h2>Devices</h2>
        <ul>
          {devices.map((device, index) => (
            <li key={index}>{device.name}</li>
          ))}
        </ul>
      </Layout>
    )
  }

  componentDidMount() {
    const { dispatch } = this.props

    dispatch(fetchDevicesIfNeeded())
  }

}

const mapState = (state) => {
  const { devices } = state.device

  return { devices }
}

export default connect(
  mapState
)(Devices)

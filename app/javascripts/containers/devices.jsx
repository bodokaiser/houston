import React, { Component } from 'react'
import { connect } from 'react-redux'

import Card from '../components/card'
import Tabs from '../components/tabs'
import Container from '../components/container'

import {
  fetchDevicesIfNeeded,
  requestDevices,
  receiveDevices
} from '../actions/device'


class Devices extends Component {

  render() {
    const { devices } = this.props

    const links = [
      { name: 'Constant', active: true },
      { name: 'Sweep' }
    ]

    return (
      <Container>
        <div className="row">
          <div className="col-sm">
            <h2 className="text-center">Devices</h2>
          </div>
        </div>
        {devices.map((device, index) => (
          <div className="row mt-3" key={index}>
            <div className="col-sm">
              <Card key={index} title={device.name}>
                <Tabs links={links} />
              </Card>
            </div>
          </div>
        ))}
      </Container>
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

import React, {
  Component,
  Fragment
} from 'react'
import {connect} from 'react-redux'

import {DetailedCard} from '../components/card'
import {InlineForm} from '../components/form'
import {NavTabs} from '../components/nav'
import {Container} from '../components/layout'

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
          {devices.map((device, index) => (
            <div className="col-sm-6">
              <DetailedCard key={index} title={device.name}>
                <Fragment>
                  <NavTabs links={links} />
                  <InlineForm />
                </Fragment>
              </DetailedCard>
            </div>
          ))}
        </div>
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

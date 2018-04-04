import React, {
  Component,
  Fragment
} from 'react'
import {connect} from 'react-redux'

import {DetailedCard} from '../components/card'
import {InlineForm} from '../components/form'
import {NavTabs} from '../components/nav'

import {
  selectDeviceTab,
  fetchDevicesIfNeeded,
  requestDevices,
  receiveDevices
} from '../actions/device'


class Device extends Component {

  constructor(props) {
    super(props)

    this.handleTabClick = this.handleTabClick.bind(this)
  }

  handleTabClick() {
    this.props.dispatch(selectDeviceTab())
  }

  render() {
    const { device } = this.props

    const links = [
      { name: 'Constant', active: true },
      { name: 'Sweep' }
    ]

    return (
      <DetailedCard title={device.name}>
        <Fragment>
          <NavTabs links={links} onClick={this.handleTabClick} />
          <InlineForm />
        </Fragment>
      </DetailedCard>
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
)(Device)

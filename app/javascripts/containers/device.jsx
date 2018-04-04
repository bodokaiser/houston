import React, {
  Component,
  Fragment
} from 'react'
import {connect} from 'react-redux'

import {DetailedCard} from '../components/card'
import {InlineForm} from '../components/form'
import {NavTabs} from '../components/nav'

import {
  updateDeviceMode
} from '../actions/local'
import {
  fetchDevicesIfNeeded
} from '../actions/remote'


class Device extends Component {

  constructor(props) {
    super(props)

    this.handleTabClick = this.handleTabClick.bind(this)
  }

  handleTabClick(mode, e) {
    e.preventDefault()

    this.props.dispatch(updateDeviceMode(this.props.device, mode))
  }

  render() {
    const { device, links } = this.props

    links.forEach(link => {
      if (link.name == device.mode) link.active = true
    })

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
  const { params } = state

  return {
    links: params.modes.map(mode => ({ name: mode }))
  }
}

export default connect(
  mapState
)(Device)

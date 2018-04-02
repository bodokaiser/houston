import React from 'react'
import {connect} from 'react-redux'

const Devices = ({ devices, actions }) => (
  <section>
    <h2>Devices</h2>
  </section>
)

const mapState = (state) => (
  { devices: state.devices }
)

export default connect(mapState)(Devices)

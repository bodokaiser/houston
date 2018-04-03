import React from 'react'
import {connect} from 'react-redux'

import {Navbar} from '../components/nav'

const mapState = state => {
  const { title, links } = state.app

  return { title, links }
}

export default connect(
  mapState
)(Navbar)

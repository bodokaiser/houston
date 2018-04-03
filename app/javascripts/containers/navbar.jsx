import React from 'react'
import { connect } from 'react-redux'


import Navbar from '../components/navbar'

const mapState = state => {
  const { title } = state.app

  return { brand: title }
}

export default connect(
  mapState
)(Navbar)

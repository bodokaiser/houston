import React, { Component } from 'react'
import { connect } from 'react-redux'

import Jumbotron from '../components/jumbotron'
import Layout from '../components/layout'

class Home extends Component {

  render() {
    return (
      <Layout columns="col-sm-12">
        <Jumbotron {...this.props} />
      </Layout>
    )
  }

}

const mapState = state => {
  return {
    title: `Welcome to ${state.app.title}`,
    lead: `Your interactive access to the experiment.`
  }
}

export default connect(
  mapState
)(Home)

import React, {
  Component,
  Fragment
} from 'react'
import {connect} from 'react-redux'

import {Page} from '../components/layout'

const Card = ({ name, value }) => (
  <div className="card">
    <div className="card-body p-3 text-center">
      <div className="h1 mt-6">{ value }</div>
      <div className="text-muted mb-4">{ name }</div>
    </div>
  </div>
)

class Dashboard extends Component {

  render() {
    const { metrics } = this.props

    return (
      <Page title="Dashboard">
        <div className="row row-cards">
          {metrics.map((metric, index) => (
            <div className="col-6 col-sm-4 col-lg-2" key={index}>
              <Card {...metric} />
            </div>
          ))}
        </div>
      </Page>
    )
  }

}

const mapState = state => ({
  metrics: Object.values(state.system)
})

export default connect(
  mapState
)(Dashboard)

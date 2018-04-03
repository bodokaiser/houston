import React, {
  Component
} from 'react'
import {connect} from 'react-redux'

import {Jumbotron} from '../components/misc'
import {Container} from '../components/layout'

class Home extends Component {

  render() {
    return (
      <Container>
        <div className="row">
          <div className="col-sm">
            <Jumbotron {...this.props} />
          </div>
        </div>
      </Container>
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

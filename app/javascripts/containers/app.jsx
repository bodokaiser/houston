import React, {
  Component,
  Fragment
} from 'react'
import {connect} from 'react-redux'

import Device from './device'
import {Navbar} from '../components/nav'

class App extends Component {

  render() {
    const { devices } = this.props

    return (
      <Fragment>
        <Navbar title="Beagle" />
        <div className="container mt-5">
          <div className="row">
            {devices.map((device, index) => (
              <div className="col-sm-4" key={index}>
                <Device key={index} device={device} />
              </div>
            ))}
          </div>
        </div>
      </Fragment>
    )
  }

}

const mapState = state => {
  const { devices } = state

  return { devices }
}

export default connect(
  mapState
)(App)

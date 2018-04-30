import React, {
  Component,
  Fragment
} from 'react'
import {
  Route,
  Switch,
  withRouter
} from 'react-router-dom'
import {connect} from 'react-redux'

import Device from './device'

import {
  Header,
  Page,
  PageNotFound
} from '../components/layout'
import {
  SimpleCard
} from '../components/card'

import {fetchDevicesLazy} from '../actions/device'

class App extends Component {

  componentDidMount() {
    //this.props.dispatch(fetchDevicesLazy())
  }

  render() {
    const { devices, metrics } = this.props

    return (
      <Fragment>
        <Header />
        <div className="container mt-5">
          <Switch>
            <Route path="/" exact render={() => (
              <Page title="Dashboard">
                <div className="row row-cards">
                  {metrics.map((metric, index) => (
                    <div className="col-6 col-sm-4 col-lg-2" key={index}>
                      <SimpleCard {...metric} />
                    </div>
                  ))}
                </div>
              </Page>
            )} />
            <Route path="/devices" render={() => (
              <Page title="Devices">
                <div className="row">
                {devices.map((device, index) => (
                  <div className="col-8 offset-2 col-sm-6 offset-sm-0 col-md-5 col-lg-4 col-xl-3" key={index}>
                    <Device key={index} device={device} />
                  </div>
                ))}
                </div>
              </Page>
            )} />
            <Route component={PageNotFound}  />
          </Switch>
        </div>
      </Fragment>
    )
  }

}

const mapState = state => ({
  devices: state.devices,
  metrics: Object.values(state.system)
})

export default withRouter(connect(
  mapState
)(App))

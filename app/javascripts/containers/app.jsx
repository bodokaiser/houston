import React, {
  Component,
  Fragment
} from 'react'

import Devices from './devices'
import Navbar from './navbar'

class App extends Component {

  render() {
    return (
      <Fragment>
        <Navbar />
        <Devices />
      </Fragment>
    )
  }

}

export default App

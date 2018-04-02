import {h} from 'preact'
import {Route, Switch} from 'react-router-dom'

import Layout from './layout'
import Navbar from './navbar'
import SynthList from './synthesizer/list'

let Index = () => (
  <h1>Hello World</h1>
)

export default () => (
  <content>
    <Navbar />
    <Layout>
      <Switch>
        <Route exact path="/" component={Index} />
        <Route path="/synthesizer" component={SynthList} />
      </Switch>
    </Layout>
  </content>
)

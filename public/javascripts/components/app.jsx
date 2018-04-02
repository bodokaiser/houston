import {h} from 'preact'
import {Route, Switch} from 'react-router-dom'

import Layout from './layout'
import Navbar from './navbar'
import SynthList from './synth/list'

let Index = () => (
  <h1>Hello World</h1>
)

export default () => (
  <content>
    <Navbar />
    <Layout>
      <Switch>
        <Route exact path="/" component={Index} />
        <Route path="/synth" component={SynthList} />
      </Switch>
    </Layout>
  </content>
)

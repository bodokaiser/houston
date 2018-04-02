import {h, render} from 'preact'
import Router from 'preact-router'

import App from './components/app'
import SynthList from './components/synth/list'

let content = document.querySelector('content')

render(
  <Router>
    <App path="/" />
    <SynthList path="/synth" /> 
  </Router>,
  content, content.firstElementChild)

if (module.hot) module.hot.accept()

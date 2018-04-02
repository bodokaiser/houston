import {h, render} from 'preact'
import {BrowserRouter as Router} from 'react-router-dom'

import App from './components/app'

let content = document.querySelector('content')

render(
  <Router>
    <App />
  </Router>, content, content.lastChild)

if (module.hot) module.hot.accept()

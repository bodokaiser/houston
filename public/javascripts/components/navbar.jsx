import {h} from 'preact'
import {Link} from 'react-router-dom'

export default () => (
  <nav class="navbar navbar-dark bg-dark sticky-top">
    <a class="navbar-brand" href="/">Beagle</a>
    <ul class="navbar-nav">
      <li class="nav-item">
        <Link class="nav-link" to="/synth">Synthesizers</Link>
      </li>
    </ul>
  </nav>
)

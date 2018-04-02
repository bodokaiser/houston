import {h} from 'preact'
import {Link} from 'react-router'

export default () => (
  <nav class="navbar navbar-dark bg-dark sticky-top">
    <a class="navbar-brand" href="/">Beagle</a>
    <ul class="navbar-nav">
      <li class="nav-item">
        <Link class="nav-link" to="/synthesizer">Synthesizer</Link>
      </li>
    </ul>
  </nav>
)

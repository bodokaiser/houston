import {h} from 'preact'

export default ({ children }) => (
  <div class="container mt-5">
      <div class="row">
        <div class="col-sm">
          <div class="jumbotron">
            <h1 class="display-4">Welcome to Beagle</h1>
            <p class="lead">Foobar</p>
          </div>
        </div>
        { children }
      </div>
  </div>
)

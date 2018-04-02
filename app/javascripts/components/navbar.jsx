import React from 'react'
import {NavLink} from 'react-router-dom'

export default () => (
  <nav className="navbar navbar-dark bg-dark sticky-top">
    <NavLink className="navbar-brand" to="/">Beagle</NavLink>
    <ul className="navbar-nav">
      <li className="nav-item">
        <NavLink className="nav-link" to="/devices">Devices</NavLink>
      </li>
    </ul>
  </nav>
)

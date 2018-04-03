import React from 'react'
import { NavLink } from 'react-router-dom'

const Navbar = ({ brand }) => (
  <nav className="navbar navbar-dark bg-dark sticky-top">
    <NavLink className="navbar-brand" to="/">{ brand }</NavLink>
    <ul className="navbar-nav">
      <li className="nav-item">
        <NavLink className="nav-link" to="/devices">Devices</NavLink>
      </li>
    </ul>
  </nav>
)

export default Navbar

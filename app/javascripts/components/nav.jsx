import React from 'react'
import {NavLink} from 'react-router-dom'

export const Navbar = ({ title, links }) => (
  <nav className="navbar navbar-dark bg-dark sticky-top">
    <NavLink className="navbar-brand" to="/">{ title }</NavLink>
    <ul className="navbar-nav">
      {links.map((link, index) =>
      <li className="nav-item" key={index}>
        <NavLink className="nav-link" to={link.path}>{link.name}</NavLink>
      </li>
      )}
    </ul>
  </nav>
)

export const NavTabs = ({ links }) => (
  <ul className="nav nav-tabs mb-3">
    {links.map((link, index) => (
      <li className="nav-item" key={index}>
        <a className={`nav-link ${link.active ? 'active' : ''}`} href="#">
          {link.name}
        </a>
      </li>
    ))}
  </ul>
)

import React from 'react'

const Tabs = ({ links }) => (
  <ul className="nav nav-tabs">
    {links.map((link, index) => (
      <li className="nav-item" key={index}>
        <a className={`nav-link ${link.active ? 'active' : ''}`} href="#">
          {link.name}
        </a>
      </li>
    ))}
  </ul>
)

export default Tabs

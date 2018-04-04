import React from 'react'

export const Navbar = ({ title, links }) => {
  links = links || []

  return (
    <nav className="navbar navbar-dark bg-dark sticky-top">
      <a className="navbar-brand" href="/">{ title }</a>
      <ul className="navbar-nav">
        {links.map((link, index) =>
        <li className="nav-item" key={index}>
          <a className="nav-link" href={link.href}>{ link.name }</a>
        </li>
        )}
      </ul>
    </nav>
  )
}


export const NavTabs = ({ links, onClick }) => (
  <ul className="nav nav-tabs mb-3">
    {links.map((link, index) => (
      <li className="nav-item" key={index}>
        <a className={`nav-link ${link.active ? 'active' : ''}`}
           onClick={e => onClick(link.name, e)} href={link.href || '/'}>
          {link.name}
        </a>
      </li>
    ))}
  </ul>
)

import React from 'react'

export const Navbar = ({ title, children }) => {
  return (
    <nav className="navbar sticky-top">
      <a className="navbar-brand" href="/">{ title }</a>
      { children }
    </nav>
  )
}


export const NavTabs = ({ links, onClick }) => (
  <ul className="nav nav-tabs mb-3">
    {links.map((link, index) => (
      <li className="nav-item" key={index}>
        <a className={`nav-link ${link.active ? 'active' : ''}`}
           href={link.href || '/'}
           onClick={e => {
             e.preventDefault()

             onClick(link.name, e)
           }}>
          {link.name}
        </a>
      </li>
    ))}
  </ul>
)

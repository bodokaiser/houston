import React from 'react'
import {NavLink} from 'react-router-dom'

export const Page = ({ children, title }) => (
  <section>
    <div className="page-header">
      <h1 className="page-title">{ title }</h1>
    </div>
    { children }
  </section>
)

export const PageError = ({ code }) => (
  <section className="page-content mt-9">
    <div className="container text-center">
      <div className="display-1 text-muted mb-5">
        <i className="si si-exclamation"></i> { code }
      </div>
      <h1 className="h2 mb-3">
        Oops.. You just found an error page..
      </h1>
      <p className="h4 text-muted font-weight-normal mb-7">
        We are sorry but our service is currently not available…
      </p>
    </div>
  </section>
)

export const PageNotFound = () => (
  <PageError code={404} />
)

export const Header = () => (
  <header className="header collapse d-lg-flex p-0">
    <div className="container">
      <div className="row align-items-center">
        <div className="col-lg order-lg-first">
          <ul className="nav nav-tabs border-0 flex-column flex-lg-row">
            <li className="nav-item">
              <NavLink className="nav-link" activeClassName="active" exact to="/">
                <i className="fe fe-home"></i> Home
              </NavLink>
            </li>
            <li className="nav-item">
              <NavLink className="nav-link" activeClassName="active" to="/devices">
                <i className="fe fe-cpu"></i> Devices
              </NavLink>
            </li>
            <li className="nav-item">
              <a className="nav-link" href="https://github.com/bodokaiser/houston">
                <i className="fe fe-github"></i> Source
              </a>
            </li>
            <li className="nav-item">
              <a className="nav-link" href="https://godoc.org/github.com/bodokaiser/houston">
                <i className="fe fe-file-text"></i> Documentation
              </a>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </header>
)

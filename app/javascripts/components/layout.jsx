import React from 'react'

const Layout = ({ columns, children }) => (
  <div className="container mt-5">
    <div className="row">
      <div className={columns}>
        { children }
      </div>
    </div>
  </div>
)

export default Layout

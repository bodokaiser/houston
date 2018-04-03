import React from 'react'

const Jumbotron = ({ title, lead }) => (
  <div className="container mt-5">
      <div className="row">
        <div className="col-sm">
          <div className="jumbotron">
            <h1 className="display-4">{ title }</h1>
            <p className="lead">{ lead }</p>
          </div>
        </div>
      </div>
  </div>
)

export default Jumbotron

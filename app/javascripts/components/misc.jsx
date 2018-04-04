import React from 'react'

export const Jumbotron = ({ title, lead }) => (
  <div className="jumbotron">
    <h1 className="display-4">{ title }</h1>
    <p className="lead">{ lead }</p>
  </div>
)

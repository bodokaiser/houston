import React from 'react'

const Card = ({ title, children }) => (
  <div className="card">
    <div className="card-header">
      { title }
    </div>
    <div className="card-body">
      { children }
    </div>
  </div>
)

export default Card

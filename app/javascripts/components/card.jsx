import React from 'react'

export const DetailedCard = ({ title, children }) => (
  <div className="card">
    <div className="card-header">
      { title }
    </div>
    <div className="card-body">
      { children }
    </div>
  </div>
)

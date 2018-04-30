import React from 'react'

export const SimpleCard = ({ name, value }) => (
  <div className="card">
    <div className="card-body p-3 text-center">
      <div className="h1 mt-6">{ value }</div>
      <div className="text-muted mb-4">{ name }</div>
    </div>
  </div>
)

export const CollapsableCard = ({ title, alert, children }) => (
  <div className="card">
    <div className="card-status bg-blue"></div>
    <div className="card-header">
      <h3 className="card-title">{title}</h3>
      <div className="card-options">
        <span className="card-options-collapse">
          <i className="fe fe-edit-2 mr-2"></i>
          <i className="fe fe-chevron-up"></i>
        </span>
      </div>
    </div>
    { alert &&
    <div className="card-alert alert alert-danger mb-0">
      { alert }
    </div> }
    { children }
  </div>
)

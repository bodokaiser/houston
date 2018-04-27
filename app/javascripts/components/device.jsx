import React from 'react'


export const Card = ({ children, title, description }) => (
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
    <div className="card-body">
      { children }
    </div>
    <div className="card-footer text-right">
      <div className="d-flex">
        <button type="button" className="btn btn-outline-secondary">Reset</button>
        <button type="button" className="btn btn-primary ml-auto">Update</button>
      </div>
    </div>
  </div>
)

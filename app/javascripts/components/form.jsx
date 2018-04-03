import React from 'react'

export const InlineForm = () => (
  <form className="form-inline">
    <div className="form-group">
      <label htmlFor="amplitude">Amplitude</label>
      <input className="form-control-plaintext" type="text" id="amplitude" />
    </div>
    <div className="form-group">
      <label htmlFor="frequency">Frequency</label>
      <input className="form-control-plaintext" type="text" id="frequency" />
    </div>
    <div className="form-group">
      <label htmlFor="phase">Phase</label>
      <input className="form-control-plaintext" type="text" id="phase" />
    </div>
    <button type="submit" className="btn btn-primary">
      Update
    </button>
  </form>
)

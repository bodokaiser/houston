import React from 'react'

export const InlineForm = () => (
  <form className="form-inline">
    <input className="form-control mb-2 mr-sm-3" type="text" id="amplitude"
      placeholder="1 dBm" />
    <input className="form-control mb-2 mr-sm-3" type="text" id="frequency"
      placeholder="250 MHz" />
    <button type="submit" className="btn btn-primary mb-2">
      Update
    </button>
  </form>
)

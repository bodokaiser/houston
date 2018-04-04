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

export const InlineForm2 = () => (
  <form className="form-inline">
    <input className="form-control mb-2 mr-sm-2" type="text" id="start"
      placeholder="100 MHz" />
    <input className="form-control mb-2 mr-sm-2" type="text" id="stop"
      placeholder="200 MHz" />
    <input className="form-control mb-2 mr-sm-2" type="text" id="time"
      placeholder="10 s" />
    <button type="submit" className="btn btn-primary mb-2">
      Update
    </button>
  </form>
)

import React from 'react'

export const DefaultForm = ({ children }) => (
  <form>
    { children }
  </form>
)

export const InlineForm = ({ children }) => (
  <form className="form-inline">
    { children }
  </form>
)

import {
  REQUEST_SPECS,
  RECEIVE_SPECS
} from '../actions/spec'

import {createReducer} from 'redux-create-reducer'

function receiveSpecs(state, action) {
  return action.specs
}

export default createReducer({}, {
  RECEIVE_SPECS: receiveSpecs
})

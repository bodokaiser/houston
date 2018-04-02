import {combineReducers} from 'redux'
import {
  REQUEST_DEVICES,
  RECEIVE_DEVICES,
} from '../actions/device'

const initialState = {
  devices: []
}

export default (state = initialState, action) => {
  switch (action.type) {
    case REQUEST_DEVICES:
      return state
    case RECEIVE_DEVICES:
      return { ...state, devices: action.devices }
    default:
      return state
  }
}

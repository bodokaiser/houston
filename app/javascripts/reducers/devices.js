import {createReducer} from 'redux-create-reducer'

import {
  UPDATE_DEVICE,
  REQUEST_DEVICES,
  RECEIVE_DEVICES
} from '../actions/device'

function receiveDevices(state, action) {
  return action.devices
}

function updateDevice(state, action) {
  return state.map(device => {
    if (device.id == action.device.id) {
      device = action.device
    }
    return device
  })
}

export default createReducer([], {
  UPDATE_DEVICE: updateDevice,
  RECEIVE_DEVICES: receiveDevices
})

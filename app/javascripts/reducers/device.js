import {createReducer} from 'redux-create-reducer'

import {
  UPDATE_DEVICE,
  RECEIVE_DEVICE_LIST,
  RECEIVE_DEVICE_UPDATE
} from '../actions/device'

function updateDevice(state, action) {
  return {
    ...state,
    devices: state.map(device => {
      if (device.name == action.device.name) {
        device = action.device
      }
      return device
    })
  }
}

function requestDeviceList(state, action) {
  return {
    ...state,
    isFetching: action.isFetching
  }
}

function receiveDeviceList(state, action) {
  return {
    ...state,
    devices: action.devices.map(device => {
      device.frequency /= 1e6
      device.amplitude *= 100

      return device
    }),
    isFetching: action.isFetching
  }
}

function requestDeviceUpdate(state, action) {
  return {
    ...state,
    device: action.device,
    isUpdating: action.isUpdating
  }
}

function receiveDeviceUpdate(state, action) {
  return {
    ...state,
    device: action.device,
    isUpdating: action.isUpdating
  }
}

export default createReducer({
  devices: [
    {
      id: 0,
      name: "Champion Board",
      amplitude: {
        "value": 1.0,
      },
      frequency: {
        "value": 1e6,
      },
      phase: {
        "value": 0,
      }
    },
    {
      id: 1,
      name: "Zinker Board",
      amplitude: {
        "value": 0.5,
      },
      frequency: {
        "value": 5e6,
      },
      phase: {
        "value": 0,
      }
    }
  ]
}, {
  UPDATE_DEVICE: updateDevice,
  REQUEST_DEVICE_LIST: requestDeviceList,
  RECEIVE_DEVICE_LIST: receiveDeviceList,
  REQUEST_DEVICE_UPDATE: requestDeviceUpdate,
  RECEIVE_DEVICE_UPDATE: receiveDeviceUpdate
})

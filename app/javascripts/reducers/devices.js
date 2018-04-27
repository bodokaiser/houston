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

export default createReducer([
  {
    id: 0,
    name: "Champion Board",
    description: "Direct Digital Synthesizer #0",
    amplitude: {
      mode: 'playback',
      const: {
        value: '100%'
      },
      sweep: {
        start: '0%',
        stop: '20%',
        duration: '1s',
        nodwells: [true, true]
      },
      playback: {
        interval: '262ns',
        trigger: false,
        duplex: true,
        data: "100%, 40%, 10%, 15%"
      },
    },
    frequency: {
      mode: 'sweep',
      const: {
        value: '10 MHz'
      },
      sweep: {
        start: '1 MHz',
        stop: '2 MHz',
        duration: '1s',
        nodwells: [true, true]
      },
      playback: {
        interval: '262ns',
        trigger: false,
        duplex: true,
        data: "1 MHz, 1 MHz, 2 MHz"
      },
    },
    phase: {
      mode: 'const',
      const: {
        value: '0 rad'
      },
      sweep: {
        start: '0 rad',
        stop: 'pi/2 rad',
        duration: '1s',
        nodwells: [true, true]
      },
      playback: {
        interval: '262ns',
        trigger: false,
        duplex: true,
        data: "0, 1.3, 2.0"
      },
    }
  },
  {
    id: 1,
    name: "Bad Board",
    description: "Direct Digital Synthesizer #1",
    amplitude: {
      mode: 'playback',
      const: {
        value: '100%'
      },
      sweep: {
        start: '0%',
        stop: '100%',
        duration: '1s',
        nodwells: [true, true]
      },
      playback: {
        interval: '262ns',
        trigger: false,
        duplex: true,
        data: "100%, 40%, 40%, 40%"
      },
    },
    frequency: {
      mode: 'const',
      const: {
        value: '3 MHz'
      },
      sweep: {
        start: '10 MHz',
        stop: '20 MHz',
        duration: '1s',
        nodwells: [true, true]
      },
      playback: {
        interval: '262ns',
        trigger: false,
        duplex: true,
        data: "5 MHz, 2 MHz, 7 MHz"
      },
    },
    phase: {
      mode: 'const',
      const: {
        value: '0 rad'
      },
      sweep: {
        start: '0 rad',
        stop: 'pi/2 rad',
        duration: '1s',
        nodwells: [true, true]
      },
      playback: {
        interval: '262ns',
        trigger: false,
        duplex: true,
        data: "0 rad, 5 rad"
      },
    }
  }
], {
  UPDATE_DEVICE: updateDevice,
  REQUEST_DEVICE_LIST: requestDeviceList,
  RECEIVE_DEVICE_LIST: receiveDeviceList,
  REQUEST_DEVICE_UPDATE: requestDeviceUpdate,
  RECEIVE_DEVICE_UPDATE: receiveDeviceUpdate
})

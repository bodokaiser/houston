import {
  REQUEST_DEVICES,
  RECEIVE_DEVICES,
} from '../actions/remote'
import {
  UPDATE_DEVICE
} from '../actions/local'

const initialState = [
  {
    id: 0,
    name: 'Signal Generator 0',
    mode: 'Single Tone',
    params: {
      singleTone: {
        amplitude: 0,
        frequency: 250e6,
      },
      linearSweep: {
        startFrequency: 100e6,
        stopFrequency: 200e6,
        timerInterval: 1,
      }
    }
  },
  {
    id: 1,
    name: 'Signal Generator 1',
    mode: 'Linear Sweep',
    params: {
      singleTone: {
        amplitude: -80,
        frequency: 300e6,
      },
      linearSweep: {
        startFrequency: 10e6,
        stopFrequency: 20e6,
        timerInterval: .5,
      }
    }
  },
]

export default (state = initialState, action) => {
  switch (action.type) {
    case UPDATE_DEVICE:
      return state.map(device => {
        if (device.id == action.device.id) {
          device = action.device
        }
        return device
      })
    default:
      return state
  }
}

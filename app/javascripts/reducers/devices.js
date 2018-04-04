import {
  REQUEST_DEVICES,
  RECEIVE_DEVICES,
} from '../actions/remote'
import {
  UPDATE_DEVICE
} from '../actions/local'

const initialState = [
  {
    name: 'Signal Generator 0',
    mode: 'Single Tone'
  },
  {
    name: 'Signal Generator 1',
    mode: 'Linear Sweep'
  },
]

export default (state = initialState, action) => {
  switch (action.type) {
    case UPDATE_DEVICE:
      return state.map(device => {
        if (device.name == action.device.name) {
          device = action.device
        }
        return device
      })
    default:
      return state
  }
}

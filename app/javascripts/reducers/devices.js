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
    mode: 'Single Tone'
  },
  {
    id: 1,
    name: 'Signal Generator 1',
    mode: 'Linear Sweep'
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

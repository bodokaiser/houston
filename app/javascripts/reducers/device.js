import {
  REQUEST_DEVICES,
  RECEIVE_DEVICES,
} from '../actions/device'

const initialState = {
  devices: [
    { name: 'Signal Generator 0', mode: 'const', type: 'AD9910' },
    { name: 'Signal Generator 1', mode: 'sweep', type: 'AD9910' }
  ],
  options: {
    frequency: {
      min: 0,
      max: 400e6
    },
    amplitude: {
      min: -85,
      max: 0
    }
  }
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

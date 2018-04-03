import {
  REQUEST_DEVICES,
  RECEIVE_DEVICES,
} from '../actions/device'

const initialState = {
  devices: [
    { name: 'Signal Generator 0', mode: 'const' },
    { name: 'Signal Generator 1', mode: 'sweep' }
  ]
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

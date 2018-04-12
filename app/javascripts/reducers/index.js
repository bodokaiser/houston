import {combineReducers} from 'redux'

import devicesReducers from './devices'

export default combineReducers({
  devices: devicesReducers
})

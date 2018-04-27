import {combineReducers} from 'redux'

import devicesReducers from './devices'
import systemReducers from './system'

export default combineReducers({
  devices: devicesReducers,
  system: systemReducers
})

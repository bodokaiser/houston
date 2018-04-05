import {combineReducers} from 'redux'

import specsReducers from './specs'
import devicesReducers from './devices'

export default combineReducers({
  specs: specsReducers,
  devices: devicesReducers
})

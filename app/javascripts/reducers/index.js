import {combineReducers} from 'redux'

import deviceReducers from './device'

export default combineReducers({
  device: deviceReducers
})

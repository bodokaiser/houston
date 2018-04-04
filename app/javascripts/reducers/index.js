import {combineReducers} from 'redux'

import deviceReducer from './device'

export default combineReducers({
  device: deviceReducer
})

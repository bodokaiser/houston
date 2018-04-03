import { combineReducers } from 'redux'

import appReducer from './app'
import deviceReducer from './device'


export default combineReducers({
  app: appReducer,
  device: deviceReducer
})

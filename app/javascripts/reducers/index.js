import {combineReducers} from 'redux'

import devicesReducer from './devices'
import paramsReducer from './params'

export default combineReducers({
  devices: devicesReducer,
  params: paramsReducer
})

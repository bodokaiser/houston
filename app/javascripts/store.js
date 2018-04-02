import {
  createStore,
  applyMiddleware
} from 'redux'
import promiseMiddleware from 'redux-promise-middleware'

import reducers from './reducers'

export function configureStore(state) {
  return createStore(reducers, state, applyMiddleware(promiseMiddleware))
}

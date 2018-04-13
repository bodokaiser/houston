import {
  compose,
  createStore,
  applyMiddleware
} from 'redux'
import thunkMiddleware from 'redux-thunk'
import loggerMiddleware from 'redux-logger'

import rootReducer from './reducers'

const middleware = [
  thunkMiddleware,
  loggerMiddleware
]

const setupStore = (state) => {
  const store = createStore(rootReducer, state, applyMiddleware(...middleware))

  if (module.hot) module.hot.accept('./reducers', () => {
    store.replaceReducer(require('./reducers').default)
  })

  return store
}

export default setupStore()

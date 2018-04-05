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

  if (process.env.NODE_ENV == 'production') {
    store.subscribe(() => (
      localStorage.setItem('state', JSON.stringify(store.getState()))
    ))
  }

  return store
}

const restoreState = () => {
  const stateString = localStorage.getItem('state')

  return (stateString) ? JSON.parse(stateString) : {}
}

const store = setupStore(restoreState())


export default store

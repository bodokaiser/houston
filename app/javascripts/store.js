import {
  createStore,
  applyMiddleware
} from 'redux'
import thunkMiddleware from 'redux-thunk'

import reducers from './reducers'

export function configureStore(state) {
  const store = createStore(reducers, state,
    applyMiddleware(thunkMiddleware))

  if (module.hot) module.hot.accept('./reducers', () => {
    store.replaceReducer(require('./reducers').default)
  })

  return store
}

export const REQUEST_SPECS = 'REQUEST_SPECS'
export const RECEIVE_SPECS = 'RECEIVE_SPECS'

export function hasSpecs(state) {
  return !!state.specs && Object.keys(state.specs).length > 0
}

export function requestSpecs() {
  return { type: REQUEST_SPECS }
}

export function receiveSpecs(specs) {
  return { type: RECEIVE_SPECS, specs }
}

export function fetchSpecs() {
  console.log('fetch specs')

  return dispatch => {
    dispatch(requestSpecs())
    console.log('requested specs')

    return fetch(`${process.env.RESOURCE}/specs`)
      .then(resp => resp.json())
      .then(json => dispatch(receiveSpecs(json)))
  }
}

export function fetchSpecsLazy() {
  return (dispatch, getState) => {
    if (hasSpecs(getState())) return

    return dispatch(fetchSpecs())
  }
}

export const REQUEST_DEVICES = 'REQUEST_DEVICES'
export const RECEIVE_DEVICES = 'RECEIVE_DEVICES'

export function requestDevices() {
  return { type: REQUEST_DEVICES }
}

export function receiveDevices(devices) {
  return { type: RECEIVE_DEVICES, devices }
}

export function fetchDevices() {
  return dispatch => {
    dispatch(requestDevices())

    return fetch('/devices')
      .then(resp => resp.json())
      .then(json => receiveDevices(json))
  }
}

export function shouldFetchDevices(state) {
  return false
}

export function fetchDevicesIfNeeded() {
  return (dispatch, getState) => {
    if (shouldFetchDevices(getState())) {
      return dispatch(fetchDevices())
    }
  }
}

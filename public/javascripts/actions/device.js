export const REQUEST_SYNTHS = 'REQUEST_DEVICES'
export const RECEIVE_SYNTHS = 'RECEIVE_DEVICES'

export function requestDevices() {
  return { type: REQUEST_DEVICES }
}

export function receiveDevices(devices) {
  return { type: RECEIVE_DEVICES, devices }
}

export function fetchSynths() {
  return dispatch => {
    dispatch(requestDevices())

    return fetch('https://www.reddit.com/r/python.json')
      .then(resp => resp.json())
      .then(json => dispatch(receiveDevices(json)))
  }
}

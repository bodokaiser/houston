export const REQUEST_SYNTHS = 'REQUEST_DEVICES'
export const RECEIVE_SYNTHS = 'RECEIVE_DEVICES'

const HOSTNAME = process.env.HOSTNAME

export function requestDevices() {
  return { type: REQUEST_DEVICES }
}

export function receiveDevices(devices) {
  return { type: RECEIVE_DEVICES, devices }
}

export function fetchSynths() {
  return dispatch => {
    dispatch(requestDevices())

    return fetch(`${HOSTNAME}/devices`)
      .then(resp => resp.json())
      .then(json => dispatch(receiveDevices(json)))
  }
}

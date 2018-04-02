export const REQUEST_DEVICES = 'REQUEST_DEVICES'
export const RECEIVE_DEVICES = 'RECEIVE_DEVICES'

const URL = process.env.URL

export function requestDevices() {
  return { type: REQUEST_DEVICES }
}

export function receiveDevices(devices) {
  return { type: RECEIVE_DEVICES, devices }
}

export function fetchDevices() {
  console.log('fetch devices')

  return dispatch => {
    dispatch(requestDevices())

    return fetch(`${URL}/devices`, { mode: 'no-cors' })
      .then(resp => resp.json())
      .then(json => receiveDevices(json))
  }
}

export const UPDATE_DEVICE = 'UPDATE_DEVICE'
export const REQUEST_DEVICES = 'REQUEST_DEVICES'
export const RECEIVE_DEVICES = 'RECEIVE_DEVICES'

export function hasDevices(state) {
  return !!state.devices && state.devices.length > 0
}

export function updateDevice(device) {
  return { type: UPDATE_DEVICE, device }
}

export function updateDeviceMode(device, mode) {
  return dispatch => {
    dispatch(updateDevice({ ...device, mode }))
  }
}
export function updateDeviceName(device, name) {
  return dispatch => {
    dispatch(updateDevice({ ...device, name }))
  }
}

export function requestDevices() {
  return { type: REQUEST_DEVICES }
}

export function receiveDevices(devices) {
  return { type: RECEIVE_DEVICES, devices }
}

export function fetchDevices() {
  return dispatch => {
    dispatch(requestDevices())

    return fetch(`${process.env.RESOURCE}/devices`)
      .then(resp => resp.json())
      .then(json => dispatch(receiveDevices(json)))
  }
}

export function fetchDevicesLazy() {
  return (dispatch, getState) => {
    if (hasDevices(getState())) return

    return dispatch(fetchDevices())
  }
}

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

export function updateDeviceSingleTone(device, singleTone) {
  device.singleTone = {...device.singleTone, ...singleTone }

  return dispatch => {
    dispatch(updateDevice(device))
  }
}

export function updateDeviceSweep(device, sweep) {
  device.sweep = { ...device.sweep, ...sweep }

  return dispatch => {
    dispatch(updateDevice(device))
  }
}

export function requestDevices() {
  return { type: REQUEST_DEVICES }
}

export function receiveDevices(devices) {
  return { type: RECEIVE_DEVICES, devices }
}

export function submitDevice(device) {
  return dispatch => {
    console.log(device)
    return fetch(`${process.env.RESOURCE}/devices/${device.id}`, {
      method: 'PUT',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ frequency: device.singleTone.frequency })
    })
  }
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

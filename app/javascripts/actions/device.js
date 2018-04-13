import axios from 'axios'

export const UPDATE_DEVICE = 'UPDATE_DEVICE'
export const REQUEST_DEVICES = 'REQUEST_DEVICES'
export const RECEIVE_DEVICES = 'RECEIVE_DEVICES'

export function hasDevices(state) {
  return !!state.devices && state.devices.length > 0
}

export function updateDevice(device) {
  return { type: UPDATE_DEVICE, device }
}

export function requestDevices() {
  return { type: REQUEST_DEVICES, isFetching: true }
}

export function receiveDevices(devices) {
  return { type: RECEIVE_DEVICES, isFetching: false, devices }
}

export function submitDevice(device) {
  return dispatch => {
    return axios.put(`/devices/dds/${device.name}`, {
        name: device.name,
        amplitude: device.amplitude,
        frequency: device.frequency,
        phase: device.phase
      }, { baseURL: process.env.RESOURCE })
      .then(resp => {
        console.log(resp)
      })
  }
}

export function fetchDevices() {
  return dispatch => {
    dispatch(requestDevices())

    return axios.get('/devices/dds', {
        baseURL: process.env.RESOURCE
      })
      .then(res => dispatch(receiveDevices(res.json)))
      .catch(err => dispatch(receiveDevices([])))
  }
}

export function fetchDevicesLazy() {
  return (dispatch, getState) => {
    if (hasDevices(getState())) return

    return dispatch(fetchDevices())
  }
}

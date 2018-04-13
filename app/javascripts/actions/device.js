import axios from 'axios'

export const UPDATE_DEVICE = 'UPDATE_DEVICE'
export const REQUEST_DEVICE_LIST = 'REQUEST_DEVICE_LIST'
export const RECEIVE_DEVICE_LIST = 'RECEIVE_DEVICE_LIST'
export const REQUEST_DEVICE_UPDATE = 'REQUEST_DEVICE_UPDATE'
export const RECEIVE_DEVICE_UPDATE = 'RECEIVE_DEVICE_UPDATE'

export function hasDevices(state) {
  return !!state.devices && state.devices.length > 0
}

export function updateDevice(device) {
  return { type: UPDATE_DEVICE, device }
}

export function requestDeviceList() {
  return { type: REQUEST_DEVICE_LIST, isFetching: true }
}

export function receiveDeviceList(devices) {
  return { type: RECEIVE_DEVICE_LIST, isFetching: false, devices }
}

export function requestDeviceUpdate(device) {
  return { type: REQUEST_DEVICE_UPDATE, isUpdating: true, device }
}

export function receiveDeviceUpdate() {
  return { type: RECEIVE_DEVICE_UPDATE, isUpdating: false }
}

export function submitDevice(device) {
  return dispatch => {
    dispatch(requestDeviceUpdate(device))

    return axios.put(`/devices/dds/${device.name}`, {
        name: device.name,
        amplitude: device.amplitude / 100,
        frequency: device.frequency * 1e6,
        phase: device.phase
      }, { baseURL: process.env.RESOURCE })
      .then(res => {
        setTimeout(() => {
          dispatch(receiveDeviceUpdate())
        }, 3000)
      })
  }
}

export function fetchDevices() {
  return dispatch => {
    dispatch(requestDeviceList())

    return axios.get('/devices/dds', {
        baseURL: process.env.RESOURCE
      })
      .then(res => dispatch(receiveDeviceList(res.data)))
      .catch(err => dispatch(receiveDeviceList([])))
  }
}

export function fetchDevicesLazy() {
  return (dispatch, getState) => {
    if (hasDevices(getState())) return

    return dispatch(fetchDevices())
  }
}

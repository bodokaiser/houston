export const UPDATE_DEVICE = 'UPDATE_DEVICE'

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

export function updateDevice(device) {
  return { type: UPDATE_DEVICE, device }
}

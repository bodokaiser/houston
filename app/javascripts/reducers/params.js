const initialState = {
  modes: [
    'Single Tone',
    'Linear Sweep'
  ],
  frequency: {
    min: 0,
    max: 400e6
  },
  amplitude: {
    min: -85,
    max: 0
  }
}

export default (state = initialState, action) => {
  return state
}

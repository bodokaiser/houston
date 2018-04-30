import {
  isEmpty,
  isDecimal
} from 'validator'
import convert from 'convert-units'

export function required() {
  return function(value) {
    return !isEmpty(value)
  }
}

export function quantity(quantity) {
  return function(value) {
    if (typeof value != 'string') return false

    var [value, abbr] = value.trim().split(' ')
    var unit = convert().getUnit(abbr)

    if (abbr == '%') {
      return isDecimal(value) && quantity == 'relative'
    }

    return isDecimal(value) && unit && unit.measure == quantity
  }
}

export function quantities(measure) {
  const fn = quantity(measure)

  return function(value) {
    return value.split(',').every(fn)
  }
}

export function modes() {
  return function(device) {
    var [nc, ns, np] = [0, 0, 0]

    var modes = [
      device.amplitude.mode,
      device.frequency.mode,
      device.phase.mode
    ]
    modes.forEach(m => {
      if (m == 'const') nc++
      if (m == 'sweep') ns++
      if (m == 'playback') np++
    })

    return np < 2 && ns < 2
  }
}

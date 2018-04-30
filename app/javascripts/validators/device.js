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
    var [value, abbr] = value.trim().split(' ')
    var unit = convert().getUnit(abbr)

    if (unit == '%') {
      return isDecimal(value) && quantity == 'relative'
    }

    return isDecimal(value) && unit && unit.measure == quantity
  }
}

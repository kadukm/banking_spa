export function prepareInn(value) {
    const newValue = value.replace(/[^0-9]/, '')
    return newValue.substring(0, 10)
}

export function prepareBik(value) {
    const newValue = value.replace(/[^0-9]/, '')
    return newValue.substring(0, 9)
}

export function prepareAccountNumber(value) {
    const newValue = value.replace(/[^0-9]/, '')
    return newValue.substring(0, 20)
}

export function prepareAmount(value) {
    const newValue = value.replace(/^0+/, '').replace(/[^0-9]/, '')
    if (newValue === '')
        return 0
    return parseInt(newValue)
}

export function preparePhone(value) {
    const newValue = value.replace(/[^0-9+]/, '')
    return newValue.substring(0, 12)
}

function stringIsInteger(value) {
    for (var i = 0; i < value.length; i++) {
        let char = value.charAt(i)
        if (isNaN(char))
            return false
    }
    return true
}

export function isInnOk(value) {
    return stringIsInteger(value) && value.length == 10
}

export function isBikOk(value) {
    return stringIsInteger(value) && value.length == 9
}

export function isAccountNumberOk(value) {
    return stringIsInteger(value) && value.length == 20
}

export function isForWhatOk(value) {
    return value.includes('без НДС') || value.includes('НДС 10%') || value.includes('НДС 18%')
}

export function isAmountOk(value) {
    return value >= 1000 && value <= 75000
}

export function isPhoneOk(value) {
    return /\+[0-9]{11}/.test(value)
}

export function isEmailOk(value) {
    return /\w+@\w+\.\w+/.test(value)
}

const ndsValues = [/НДС 18%/g, /НДС 10%/g, /без НДС/g]

export const csrfTokenName = "X-CSRF-TOKEN"

export function getCookie(cname) {
    var name = cname + "=";
    var ca = document.cookie.split(';');
    for (var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

export function urlEncodeObject(obj) {
    var str = [];
    for (var p in obj)
        if (obj.hasOwnProperty(p)) {
            str.push(encodeURIComponent(p) + "=" + encodeURIComponent(obj[p]));
        }
    return str.join("&");
}

export function handleNds(currentNdsValue, inputValue) {
    var newInputValue = inputValue
    for (let ndsValue of ndsValues) {
        newInputValue = newInputValue.replace(ndsValue, "")
    }
    newInputValue = newInputValue.replace(/ $/, "")
    newInputValue = newInputValue + ' ' + currentNdsValue
    return newInputValue.replace(/^ +/, "")
}

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
        return ''
    return parseInt(newValue)
}

export function preparePhone(value) {
    const newValue = value.replace(/[^0-9+]/, '')
    return newValue.substring(0, 12)
}

export function prepareCardNumber(value) {
    const newValue = value.replace(/[^0-9]/, '')
    return newValue.substring(0, 16)
}

export function prepareCardExpires(value) {
    const newValue = value.replace(/[^0-9/]/, '')
    return newValue.substring(0, 5)
}

export function prepareCardCvc(value) {
    const newValue = value.replace(/[^0-9]/, '')
    return newValue.substring(0, 3)
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

export function isCardNumberOk(value) {
    return stringIsInteger(value) && value.length == 16
}

export function isCardExpiresOk(value) {
    return /^(0[1-9]|1[0-2])\/(19|[23]\d|2[0-5])$/.test(value)
}

export function isCardCvcOk(value) {
    return stringIsInteger(value) && value.length == 3
}

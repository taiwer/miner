import Cookies from 'js-cookie'

const TokenKey = 'jwt-token'
const TokenExpireKey = 'token-expire'

export function getToken () {
  const token = Cookies.get(TokenKey)
  return token
}

export function setToken (token) {
  var inFifteenMinutes = new Date(new Date().getTime() + 6 * 60 * 60 * 1000)
  return Cookies.set(TokenKey, token, { expires: inFifteenMinutes })
}

export function removeToken () {
  return Cookies.remove(TokenKey)
}
export function getTokenExpire () {
  const tokenExpire = Cookies.get(TokenExpireKey)
  return tokenExpire
}

export function setTokenExpire (tokenExpire) {
  var inFifteenMinutes = new Date(new Date().getTime() + 6 * 60 * 60 * 1000)
  return Cookies.set(TokenExpireKey, tokenExpire, { expires: inFifteenMinutes })
}

export function removeTokenExpire () {
  return Cookies.remove(TokenExpireKey)
}

export function sleep (time) {
  return new Promise((resolve) => setTimeout(resolve, time))
}

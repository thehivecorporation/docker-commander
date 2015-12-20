/** Actions.jsx */

export const REFRESH_REQUEST = 'REFRESH_REQUEST'
export const REFRESH_SUCCESS = 'REFRESH_SUCCESS'
export const REFRESH_FAILED = 'REFRESH_FAILED'
export const OPEN_SOCKET = 'OPEN_SOCKET'
export const CLOSE_SOCKET = 'CLOSE_SOCKET'

export function refreshRequest(){
  return {
    type:REFRESH_REQUEST
  }
}

export function refreshSuccess(json){
  return {
    type:REFRESH_SUCCESS,
    data:json
  }
}

export function refreshFailed(json){
  return {
    type:REFRESH_FAILED,
    data:json
  }
}

export function openSocket(){
  return {
    type: OPEN_SOCKET,
  }
}

export function closeSocket(){
  return {
    type: CLOSE_SOCKET
  }
}

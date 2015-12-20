import store from './store.jsx'
import {REFRESH_REQUEST, REFRESH_FAILED, REFRESH_SUCCESS } from './actions.jsx';
import { refreshRequest, refreshFailed, refreshSuccess } from './actions.jsx';

const url = 'ws://localhost:8000/ws';
let socket = new WebSocket(url);

socket.onopen = ()  => {
  store.dispatch(refreshRequest())
};
socket.onclose = () => { console.log("Socket closed..."); };

socket.onmessage = msg => {
  let json = JSON.parse(msg.data);
  switch (json.Action) {
    case 'cluster':
      store.dispatch(refreshSuccess(json.Response))
      break;
    default:
      console.log("action not recognized", json.Action)
      break;
  }
}

export default socket;

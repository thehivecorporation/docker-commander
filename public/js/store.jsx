import {REFRESH_REQUEST, REFRESH_FAILED, REFRESH_SUCCESS } from './actions.jsx';
import { refreshRequest, refreshFailed, refreshSuccess } from './actions.jsx';
import { createStore } from 'redux';
import socket from './socket.jsx';

let counter = function(state = {}, action){
  switch(action.type){
    case 'REFRESH_REQUEST':
      console.log("REFRESH_REQUEST:");
      socket.send(JSON.stringify({action:'cluster'}));
      return state
    case 'REFRESH_FAILED':
      console.error("Store refresh failed");
      return state;
    case 'REFRESH_SUCCESS':
      console.log('Refresh success');
      return action.data;
    default:
      return state;
  }
}

let store = createStore(counter);

export default store

import React from 'react';

class Agent extends React.Component {
  render(){
    return (
      <div>
        <h2>Agent</h2>
        <div>IP:</div><div>{this.props.agent.IP}</div>
      </div>
  );
  }
}

export default Agent

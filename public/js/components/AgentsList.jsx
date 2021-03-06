import React from 'react';
import Agent from './Agent.jsx';

class AgentsList extends React.Component {
    render(){
      let agents = this.props.agents.map(a => {
        return <Agent style={this.props.style} agent={a} key={a.IP} />
      });

      return(
        <div>{agents}</div>
      )
    }
}

export default AgentsList

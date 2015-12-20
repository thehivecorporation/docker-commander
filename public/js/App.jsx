import React from 'react';
import { connect } from 'react-redux';
import Cluster from './components/Cluster.jsx'
import AgentsList from './components/AgentsList.jsx'

class App extends React.Component {
  render() {
    console.log("Props:", this.props);

    if(!this.props.Cluster){
      return <div>Loading data from server...</div>;

    } else {
      return (
        <div>
          <Cluster cluster={this.props.Cluster} />
          <AgentsList agents={this.props.Agents} />
        </div>
      );
    }
  }
}

function select(state){
  return state;
}

export default connect(select)(App)

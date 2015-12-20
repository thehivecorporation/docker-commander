import React from 'react';
import { connect } from 'react-redux';
import Cluster from './components/Cluster.jsx'
import AgentsList from './components/AgentsList.jsx'
import DriverStatus from './components/DriverStatus.jsx';
const Paper = require('material-ui/lib/paper');

class App extends React.Component {
  render() {
    let style = {margin: "10px"}

    if(!this.props.Cluster){
      return (
        <Paper zDepth={2} style={style}>
          <p>Loading data from server...</p>
        </Paper>
      );

    } else {
      return (
        <div>
            <Cluster style={style} cluster={this.props.Cluster} />
            <DriverStatus style={style} driverStatus={this.props.Cluster.DriverStatus} />
            <AgentsList style={style} agents={this.props.Agents} />
        </div>
      );
    }
  }
}

function select(state){
  return state;
}

export default connect(select)(App)

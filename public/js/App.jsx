import React from 'react';
import { connect } from 'react-redux';
import Cluster from './components/Cluster.jsx'
import AgentsList from './components/AgentsList.jsx'
const Paper = require('material-ui/lib/paper');
const AppBar = require('material-ui/lib/app-bar');
const Avatar = require('material-ui/lib/avatar');

class App extends React.Component {
  render() {
    let style = {margin: "10px 5px 10px 5px"}

    if(!this.props.Cluster){
      return (
        <Paper zDepth={2} style={style}>
          <p>Loading data from server...</p>
        </Paper>
      );

    } else {
      return (
        <div>
          <AppBar
            title="Docker Commander"
            iconElementLeft={<Avatar src="img/swarm.png"></Avatar>}
          />
          <Cluster style={style} cluster={this.props.Cluster} />
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

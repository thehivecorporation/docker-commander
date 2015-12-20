import React from 'react';
import { connect } from 'react-redux';
import Cluster from './components/Cluster.jsx'
import AgentsList from './components/AgentsList.jsx'
const Paper = require('material-ui/lib/paper');
const AppBar = require('material-ui/lib/app-bar');
const Avatar = require('material-ui/lib/avatar');
const FontIcon = require('material-ui/lib/font-icon');
const IconButton = require('material-ui/lib/icon-button');
import injectTapEventPlugin from 'react-tap-event-plugin';
injectTapEventPlugin();

class App extends React.Component {
  onClickGithub(e) {
      console.log("Asdasd",e);
  }

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
            iconElementRight={
              <IconButton tooltip="GitHub" style={{padding:"0px"}} onFocus={function(e){
                  window.open("http://github.com/sayden/docker-commander");
                }}>
                <Avatar src="img/github.png"/>
              </IconButton>
              }
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

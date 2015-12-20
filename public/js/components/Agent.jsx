import React from 'react';

const Card = require('material-ui/lib/card/card');
const Avatar = require('material-ui/lib/avatar');
const CardHeader = require('material-ui/lib/card/card-header');
const TableRow = require('material-ui/lib/table/table-row');
const TableRowColumn = require('material-ui/lib/table/table-row-column');
const Table = require('material-ui/lib/table/table');
const TableHeader = require('material-ui/lib/table/table-header');
const TableHeaderColumn = require('material-ui/lib/table/table-header-column');
const TableBody = require('material-ui/lib/table/table-body');
const FontIcon = require('material-ui/lib/font-icon');
import ContainerList from './ContainerList.jsx';
import injectTapEventPlugin from 'react-tap-event-plugin';
injectTapEventPlugin();

class Agent extends React.Component {
  render(){
    let exposedPorts = this.props.agent.Containers.map(c => {
      return c.Ports.filter(p => {
        return (p.PublicPort != "" || p.PublicPort != 0);
      })
      .map(p => {
        return p.PublicPort;
      });
    }).toString();

    //Delete trailing comma in result
    exposedPorts = exposedPorts.substring(0, exposedPorts.length - 1)

    let agentInfo = "Containers: " + this.props.agent.Containers.length
      + ", Images: " + this.props.agent.Images.length
      + ", Exposed Ports: " + exposedPorts;
    return (
      <Card style={this.props.style}>
        <CardHeader
          title={"Agent " + this.props.agent.IP}
          subtitle={agentInfo}
          actAsExpander={true}
          showExpandableButton={true}
          avatar={<Avatar src="img/docker.png"></Avatar>}
        />
        <ContainerList expandable={true} containers={this.props.agent.Containers} />
      </Card>
  );
  }
}

export default Agent

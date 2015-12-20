import React from 'react';
import Agent from './Agent.jsx';
import ContainerPort from './ContainerPort.jsx';
const List = require('material-ui/lib/lists/list');
const ListDivider = require('material-ui/lib/lists/list-divider');

class ContainerPortList extends React.Component {
    render(){
      let id = this.props.containerId;

      let ports = this.props.ports.map(p => {
        return <div key={id + p.PublicPort + p.PrivatePort}><ContainerPort port={p} /></div>
      });

      return(
        <List
          subheader="List of mapped ports"
          subHeaderStyle={{height:'15px'}}
          style={{backgroundColor:'#00000000'}} //Transparent
          subheaderStyle={{height:'30px'}}
          >
          {ports}
        </List>
      )
    }
}

export default ContainerPortList

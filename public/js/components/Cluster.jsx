import React from 'react';
const Card = require('material-ui/lib/card/card');
const Avatar = require('material-ui/lib/avatar');
const CardHeader = require('material-ui/lib/card/card-header');
const CardText = require('material-ui/lib/card/card-text');

class Cluster extends React.Component {
  render(){
    let cluster = this.props.cluster;

    return(
      <Card>
        <CardHeader
          title="Cluster"
          subtitle="Information about swarm cluster"
          avatar={<Avatar src="img/swarm.png"></Avatar>}
          />
        <CardText>
          <div>
            <div>Containers:</div>
            <div>{cluster.Containers}</div>
          </div>
          <div>
            <div>Name:</div>
            <div>{cluster.Name}</div>
          </div>
        </CardText>
      </Card>
    )
  }
}

export default Cluster

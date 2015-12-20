import React from 'react';

class ContainerList extends React.Component {
  render(){
    let containers = this.props.containers.map(c => {
      return <Container container=c />
    });

    return (
      {containers}
    );
  }
}

export default ContainerList;

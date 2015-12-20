import React from 'react';
import Container from './Container.jsx';

class ContainerList extends React.Component {
  render(){
    let style = {margin:'7px'}
    let containers = this.props.containers.map(c => {
      return <Container container={c} style={style} key={c.Id} />
    });

    return (
      <div>{containers}</div>
    );
  }
}

export default ContainerList;

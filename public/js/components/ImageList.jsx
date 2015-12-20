import React from 'react';
import Image from './Image.jsx';

class ImageList extends React.Component {
  render(){
    let style = {margin:'7px'}
    let images = this.props.images.map(i => {
      return <Image image={i} style={style} key={i.Id} />
    });

    return (
      <div>{images}</div>
    );
  }
}

export default ImageList;

import React from 'react';
const ListItem = require('material-ui/lib/lists/list-item');
const List = require('material-ui/lib/lists/list');
const ListDivider = require('material-ui/lib/lists/list-divider');
var Typography = require('material-ui/lib/styles/typography');

class ContainerPort extends React.Component {
  render(){
    let port = this.props.port;
    let style = {padding:'8px'}
    let secondaryTextStyle= {
      fontSize: 14,
      lineHeight: '16px',
      height: 16,
      margin: 0,
      marginTop: 4,
      color: Typography.textLightBlack
    };

    return(
        <div>
          <ListItem
            style={{float:'left'}}
            disabled={true}>
            <div style={{padding:'0 0 10px 0'}}>
              <div>IP</div>
              <div style={secondaryTextStyle}>{port.IP}</div>
            </div>
            <div>
              <div>Type</div>
              <div style={secondaryTextStyle}>{port.Type}</div>
            </div>
          </ListItem>

          <ListItem
            style={{float:'right'}}
            disabled={true}>
            <div style={{padding:'0 0 10px 0'}}>
              <div>PublicPort</div>
              <div style={secondaryTextStyle}>{port.PublicPort}</div>
            </div>
            <div>
              <div>Private Port</div>
              <div style={secondaryTextStyle}>{port.PrivatePort}</div>
            </div>
          </ListItem>
        </div>
    );
  }
}

export default ContainerPort

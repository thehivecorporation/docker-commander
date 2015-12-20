import React from 'react';
const Card = require('material-ui/lib/card/card');
const Avatar = require('material-ui/lib/avatar');
const CardHeader = require('material-ui/lib/card/card-header');
const CardText = require('material-ui/lib/card/card-text');
const TableRow = require('material-ui/lib/table/table-row');
const TableRowColumn = require('material-ui/lib/table/table-row-column');
const Table = require('material-ui/lib/table/table');
const TableHeader = require('material-ui/lib/table/table-header');
const TableHeaderColumn = require('material-ui/lib/table/table-header-column');
const TableBody = require('material-ui/lib/table/table-body');
import Colors from 'material-ui/lib/styles/colors';
import ContainerPortList from './ContainerPortList.jsx';
import injectTapEventPlugin from 'react-tap-event-plugin';
injectTapEventPlugin();

class Container extends React.Component {
  render(){
    let container = this.props.container;
    let rows = [];
    let portInfo;

    for (var key in container) {
      if (container.hasOwnProperty(key)) {
        //Non empty info
        if (container[key] != '' && container[key] != null) {
          if( key == 'Ports' ) {
            rows.push(
              <TableRow key={container.Id + key + "_row"}>
                <TableRowColumn>{key}</TableRowColumn>
                <TableRowColumn>
                  <ContainerPortList ports={container['Ports']} containerId={container.Id} />
                </TableRowColumn>
              </TableRow>
            );

            portInfo = container.Ports.map(p => {
              return p.IP + ":" + p.PublicPort + "->" + p.PrivatePort + "/" + p.Type + "     ";
            }).toString();


          } else {
            rows.push(
              <TableRow key={container.Id + key + "_row"}>
                <TableRowColumn>{key}</TableRowColumn>
                <TableRowColumn>
                  {container[key]}
                </TableRowColumn>
              </TableRow>
            );
          }
        }
      }
    }
    return(
        <Card style={this.props.style} initiallyExpanded={false}>
          <CardHeader
            title={this.props.container.Command}
            subtitle={portInfo}
            actAsExpander={true}
            showExpandableButton={true}
            avatar={<Avatar src="img/container.jpg"></Avatar>}
            />

            <Table
              height='300px'
              fixedHeader={true}
              fixedFooter={true}
              expandable={true}
              selectable={false}>
              <TableHeader enableSelectAll={false} displaySelectAll={false}>
                <TableRow>
                  <TableHeaderColumn tooltip='The Name'>Name</TableHeaderColumn>
                  <TableHeaderColumn tooltip='The Status'>Status</TableHeaderColumn>
                </TableRow>
              </TableHeader>
              <TableBody
                displayRowCheckbox={false}
                showRowHover={true}
                >

                {rows}

              </TableBody>
            </Table>
        </Card>
    );
  }
}

export default Container;

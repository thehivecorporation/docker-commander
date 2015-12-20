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
const FontIcon = require('material-ui/lib/font-icon');
import Colors from 'material-ui/lib/styles/colors';
import DriverStatus from './DriverStatus.jsx';
import injectTapEventPlugin from 'react-tap-event-plugin';
injectTapEventPlugin();

class Cluster extends React.Component {
  render(){
    let cluster = this.props.cluster;
    let rows = [];
    let driverStatus = (<DriverStatus expandable={true} style={{margin:'7px'}}
      driverStatus={cluster.DriverStatus} />);
    delete cluster.DriverStatus

    for (var key in cluster) {
      if (cluster.hasOwnProperty(key)) {
        //Non empty info
        if (cluster[key] != '' && cluster[key] != null) {

          if(cluster[key] === true){
            rows.push(
              <TableRow key={key}>
                <TableRowColumn>{key}</TableRowColumn>
                <TableRowColumn>
                  <FontIcon className="material-icons" color={Colors.green800}>
                    check
                  </FontIcon>
                </TableRowColumn>
              </TableRow>
            );

          } else if (cluster[key] === false) {
            rows.push(
              <TableRow key={key}>
                <TableRowColumn>{key}</TableRowColumn>
                <TableRowColumn>
                  <FontIcon className="material-icons"
                    color={Colors.red800}>close</FontIcon>
                </TableRowColumn>
              </TableRow>
            );

          } else {
            rows.push(
              <TableRow key={key}>
                <TableRowColumn>{key}</TableRowColumn>
                <TableRowColumn>
                  {cluster[key]}
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
          title="Swarm"
          subtitle="Tap to get details..."
          actAsExpander={true}
          showExpandableButton={true}
          avatar={<Avatar src="img/swarm.png"></Avatar>}
        />

        <Card style={{margin:'7px'}} expandable={true} initiallyExpanded={false}>
          <CardHeader
            title="Swarm"
            actAsExpander={true}
            showExpandableButton={true}
            avatar={
              <Avatar
                icon={
                  <FontIcon className="material-icons">info_outline</FontIcon>
                }
                color={Colors.blue700}
                backgroundColor={Colors.blue100}
              />
            }
            />

          <Table
            height='400px'
            fixedHeader={true}
            fixedFooter={true}
            expandable={true}
            selectable={false}
            >
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

        {driverStatus}

      </Card>
    )
  }
}

export default Cluster

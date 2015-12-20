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
import injectTapEventPlugin from 'react-tap-event-plugin';
import Colors from 'material-ui/lib/styles/colors';
injectTapEventPlugin();

class DriverStatus extends React.Component {
  render(){
    let rows = this.props.driverStatus.map(d => {
      d[0] = d[0].replace("â”” ", "")   //TODO Research about this problem
      return(
        <TableRow key={d[0] + d[1]}>
          <TableRowColumn>{d[0]}</TableRowColumn>
          <TableRowColumn>{d[1]}</TableRowColumn>
        </TableRow>
      );
    })


    return(
      <Card style={this.props.style} initiallyExpanded={false}>
        <CardHeader
          title="Driver Status"
          subtitle="Click to expand/contract"
          avatar={
            <Avatar
              icon={
                <FontIcon className="material-icons">memory</FontIcon>
              }
              color={Colors.blue700}
              backgroundColor={Colors.blue100}
            />
          }
          actAsExpander={true}
          showExpandableButton={true}
        />


        <Table
          height='400px'
          fixedHeader={true}
          fixedFooter={true}
          expandable={true}
          selectable={false}>
          <TableHeader enableSelectAll={false} displaySelectAll={false}>
            <TableRow>
              <TableHeaderColumn tooltip='The Driver'>Driver</TableHeaderColumn>
              <TableHeaderColumn tooltip='The Info'>Info</TableHeaderColumn>
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

export default DriverStatus

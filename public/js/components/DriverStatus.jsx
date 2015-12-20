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
      return(
        <TableRow key={d[0]}>
          <TableRowColumn>{d[0]}</TableRowColumn>
          <TableRowColumn>
            {d[1]}
          </TableRowColumn>
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
          fixedHeader={false}
          fixedFooter={false}
          selectable={false}
          multiSelectable={false}
          expandable={true}
          >
          <TableHeader enableSelectAll={false} displaySelectAll={false} >
            <TableRow>
              <TableHeaderColumn tooltip='The Driver'>Driver</TableHeaderColumn>
              <TableHeaderColumn tooltip='The Status'>Status</TableHeaderColumn>
            </TableRow>
          </TableHeader>
          <TableBody
            deselectOnClickaway={false}
            showRowHover={true}
            displayRowCheckbox={false}
            >

            {rows}

          </TableBody>
        </Table>
      </Card>
    );
  }
}

export default DriverStatus

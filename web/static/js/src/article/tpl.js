/**
 * Created by lyric on 4/9/16.
 */
var React = require('react');
var Bootstrap = require('react-bootstrap');
var Grid = Bootstrap.Grid;
var Row = Bootstrap.Row;
var Col = Bootstrap.Col;
var PageHeader = Bootstrap.PageHeader;

module.exports = React.createClass({
    render: function () {
        return (
            <Grid>
                <Row>
                    <Col md={12}>
                        <PageHeader>
                            文章管理&nbsp;
                            <small>{this.props.title}</small>
                        </PageHeader>
                        {this.props.children}
                    </Col>
                </Row>
            </Grid>
        );
    }
});
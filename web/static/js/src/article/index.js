var React = require('react');
var Bootstrap = require('react-bootstrap');
var Table = Bootstrap.Table;
var Button = Bootstrap.Button;
var Tpl = require('./tpl');

module.exports = React.createClass({
    getInitialState: function () {
        return {data: []};
    },
    componentDidMount: function () {
        $.get(this.props.url, {dataType: 'json'})
            .done(function (result) {
                if (result.status !== 0) {
                    alert(result.msg);
                    return false;
                }
                this.setState({data: result.data});
            }.bind(this));
    },
    render: function () {
        var tbodyData = this.state.data.map(function (article, i) {
            return (
                <tr>
                    <td>{i + 1}</td>
                    <td>{article.Title}</td>
                    <td>{article.Content}</td>
                    <td>{article.Author}</td>
                    <td>
                        <Button bsStyle="warning" bsSize="xs" href={'/article/update/'+article.ID}>编辑</Button>&nbsp;
                        <Button bsStyle="danger" bsSize="xs" href={'/article/delete/'+article.ID}>删除</Button>&nbsp;
                        <Button bsStyle="info" bsSize="xs" href={'/article/view/'+article.ID}>查看</Button>
                    </td>
                </tr>
            );
        });
        return (
            <Tpl title="文章列表">
                <Button bsStyle="primary" href="/article/create" style={{marginBottom:'10px'}}>新增</Button>
                <Table striped bordered condensed hover>
                    <thead>
                    <tr>
                        <th style={{width:'5%'}}>#</th>
                        <th style={{width:'20%'}}>标题</th>
                        <th>内容</th>
                        <th style={{width:'15%'}}>作者</th>
                        <th style={{width:'20%'}}>操作</th>
                    </tr>
                    </thead>
                    <tbody>
                    {tbodyData}
                    </tbody>
                </Table>
            </Tpl>
        );
    }
});
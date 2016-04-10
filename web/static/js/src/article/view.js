/**
 * Created by lyric on 4/10/16.
 */
var React = require('react');
var Bootstrap = require('react-bootstrap');
var Input = Bootstrap.Input;
var Button = Bootstrap.Button;
var ButtonToolbar = Bootstrap.ButtonToolbar;
var Tpl = require('./tpl');

module.exports = React.createClass({
    componentDidMount: function () {
        $.get('/article/list/' + this.props.id, {dataType: 'json'}).done(function (result) {
            if (result.status !== 0) {
                alert(result.msg);
                return false;
            }
            var data = result.data;
            $('#Title').val(data.Title);
            $('#Content').val(data.Content);
            $('#Author').val(data.Author);
        }.bind(this));
    },
    handleBack: function () {
        history.back();
    },
    render: function () {
        return (
            <Tpl title="查看文章">
                <form method="post">
                    <Input type="text" label="标题" placeholder="请输入标题" name="Title" id="Title" disabled/>
                    <Input type="textarea" label="内容" placeholder="请输入文章内容" name="Content" id="Content"
                           style={{height:'300px'}} disabled/>
                    <Input type="text" label="作者" placeholder="请输入作者" name="Author" id="Author" disabled/>
                    <ButtonToolbar>
                        <Button bsStyle="default" onClick={this.handleBack}>返回</Button>
                    </ButtonToolbar>
                </form>
            </Tpl>
        );
    }
});
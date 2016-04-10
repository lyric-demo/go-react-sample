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
    handleSubmit: function (e) {
        var title = $('#Title').val();
        if (!title) {
            alert("标题不能为空!");
            e.preventDefault();
            return false;
        }
    },
    handleBack: function () {
        history.back();
    },
    render: function () {
        return (
            <Tpl title="新建文章">
                <form method="post" action={this.props.url} onSubmit={this.handleSubmit}>
                    <Input type="text" label="标题" placeholder="请输入标题" name="Title" id="Title"/>
                    <Input type="textarea" label="内容" placeholder="请输入文章内容" name="Content" style={{height:'300px'}}/>
                    <Input type="text" label="作者" placeholder="请输入作者" name="Author"/>
                    <ButtonToolbar>
                        <Button type="submit" bsStyle="success">保存</Button>
                        <Button bsStyle="default" onClick={this.handleBack}>返回</Button>
                    </ButtonToolbar>
                </form>
            </Tpl>
        );
    }
});
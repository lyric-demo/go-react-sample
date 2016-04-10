/**
 * Created by lyric on 4/9/16.
 */
var React = require('react');
var ReactDOM = require('react-dom');
var Index = require('./article/index');
var Create = require('./article/create');
var Update = require('./article/update');
var Delete = require('./article/delete');
var View = require('./article/view');

var list = document.getElementById('list');
if (list) {
    ReactDOM.render(
        <Index url="/article/list"/>, list
    );
}

var create = document.getElementById('create');
if (create) {
    ReactDOM.render(
        <Create url="/article/create"/>, create
    );
}

var update = document.getElementById('update');
if (update) {
    ReactDOM.render(
        <Update id={ID}/>, update
    );
}

var del = document.getElementById('delete');
if (del) {
    ReactDOM.render(
        <Delete id={ID}/>, del
    );
}

var view = document.getElementById('view');
if (view) {
    ReactDOM.render(
        <View id={ID}/>, view
    );
}
import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import axios from "axios"

axios.defaults.withCredentials = true
axios.defaults.timeout = 100000

class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <div onClick={this.getList}>axios请求接口GET</div>
          <div onClick={this.login}>axios请求接口POST fromdata</div>

        </header>
      </div>
    );
  }
  getList() {
    console.log("开始请求!")
    axios.get("http://localhost:8080/topics")
      .then(r => {
        console.log(r)
      })
    console.log("结束请求!")
  }
  login() {
    console.log("开始login!")
    axios.post("http://localhost:8080/login", {
      email: "891177434@qq.com",
      password: "123456",
    })
      .then(r => {
        console.log(r)
      })
    console.log("结束login!")
  }
}

export default App;



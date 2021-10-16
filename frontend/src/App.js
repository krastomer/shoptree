import "./App.css";
import React, { Component } from "react";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      items: [],
    };
  }

  componentDidMount() {
    fetch("http://127.0.0.1:8080/api/v1/product/1")
      .then((res) => res.json())
      .then((json) => {
        this.setState({
          items: json,
        });
      });
  }

  render() {
    var items = this.state;
    console.log(items);

    return (
      <div className="App">
        <ul></ul>
      </div>
    );
  }
}

export default App;

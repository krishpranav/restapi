import React from 'react';
import './App.css';
import logo from './logo.svg';

class UserComponent extends React.Component {
  componentDidMount() {
    const userApi = 'http://localhost:8080/user'
    fetch (userApi)
      .then((response) => response.json())
      .then((data) => console.log('User Api Request', data))
  }

  render() {
    
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            See The Console For Api Request
            </p>
          </header>
      </div>
    )
  }
}

export default UserComponent;
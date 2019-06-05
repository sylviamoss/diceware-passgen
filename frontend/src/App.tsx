import React from 'react';
import PasswordFields, { Words } from './PasswordFields';
import './App.css';

class App extends React.Component {
  state = {
    isPasswordOpen: false,
    words: {
      plain: "",
      password: ""
    }
  }
  render() {
    return (
      <div className="App">
        <div className="App-header">
          <PasswordFields
            isOpen={this.state.isPasswordOpen}
            words={this.state.words}
          />
        </div>
      </div>
    );
  }
}

export default App;

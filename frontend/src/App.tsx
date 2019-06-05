import React from 'react';
import PasswordFields from './PasswordFields';
import LangSelector from './LangSelector';
import './App.css';

class App extends React.Component {
  state = {
    isPasswordOpen: false,
    pass: {
      words: "",
      password: ""
    }
  }

  constructor(props: any) {
    super(props)
    this.handleLangSelector = this.handleLangSelector.bind(this)
  }

  public handleLangSelector(lang: string) {
    fetch('/api/'+lang)
    .then(response => response.json())
    .then(pass => this.setState({ pass }));
    this.setState({isPasswordOpen: true})
  }

  render() {
    return (
      <div className="App">
        <div className="App-header">
          <PasswordFields
            isOpen={this.state.isPasswordOpen}
            pass={this.state.pass}
          />
          <LangSelector
            handleLangSelector={this.handleLangSelector}
          />
        </div>
      </div>
    );
  }
}

export default App;

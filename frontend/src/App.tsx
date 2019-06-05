import React from 'react';
import PasswordFields, { Words } from './PasswordFields';
import LangSelector from './LangSelector';
import './App.css';

class App extends React.Component {

  constructor(props: any) {
    super(props)

    this.handleLangSelector = this.handleLangSelector.bind(this)
  }
  state = {
    isPasswordOpen: false,
    words: {
      plain: "",
      password: ""
    }
  }

  public handleLangSelector(lang: string) {
    const words = {
      plain: "la la la " + lang,
      password: "lalala" + lang
    }
    this.setState({isPasswordOpen: true})
    this.setState({ words })  
  }

  render() {
    return (
      <div className="App">
        <div className="App-header">
          <PasswordFields
            isOpen={this.state.isPasswordOpen}
            words={this.state.words}
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

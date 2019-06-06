import React from 'react';
import PasswordFields from './PasswordFields';
import LangSelector from './LangSelector';
import { IconContext } from "react-icons";
import { FaGithub, FaDiceOne, FaDiceTwo, FaDiceThree, FaDiceFour, FaDiceFive, FaDiceSix } from 'react-icons/fa';
import './App.css';

class App extends React.Component {
  state = {
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
    fetch('/api/' + lang)
      .then(response => response.json())
      .then(pass => this.setState({ pass }));
    this.setState({ isPasswordOpen: true })
  }

  render() {
    return (
      <div className="App">
        <div className="App-header">
          <IconContext.Provider value={{ size: "3em", color: "whitesmoke", className: "dice-icon-provider" }}>
            <div className="dice-icons">
              <FaDiceOne className="dice-icon"/>
              <FaDiceTwo className="dice-icon"/>
              <FaDiceThree className="dice-icon"/>
              <FaDiceFour className="dice-icon"/>
              <FaDiceFive className="dice-icon"/>
              <FaDiceSix className="dice-icon"/>
            </div>
          </IconContext.Provider>

          <PasswordFields
            pass={this.state.pass}
          />
          <LangSelector
            handleLangSelector={this.handleLangSelector}
          />
        </div>

        <FaGithub />
      </div>
    );
  }
}

export default App;

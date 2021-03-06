import React from 'react';
import CopyToClipboard from 'react-copy-to-clipboard'
import { FaCopy } from 'react-icons/fa';
import './PasswordFields.css';

export interface Pass {
    words: string,
    password: string
}

type PasswordFieldsProps = {
    pass: Pass
}

class PasswordFields extends React.Component<PasswordFieldsProps> {
    state = {
        words: "",
        password: ""
    }

    render() {
        const { pass } = this.props

        return (
            <div className="password-fields">
                <div className="field" >
                    <input className="form-control" value={pass.words} readOnly />
                    <CopyToClipboard text={pass.words}>
                        <button type="button" className="btn btn-outline-warning">
                            <FaCopy /> copy
                        </button>
                    </CopyToClipboard>
                </div>
                <div className="field" >
                    <input className="form-control" value={pass.password} readOnly />
                    <CopyToClipboard text={pass.password}>
                        <button type="button" className="btn btn-outline-warning">
                            <FaCopy /> copy
                        </button>
                    </CopyToClipboard>
                </div>
            </div>
        )
    }
}

export default PasswordFields;
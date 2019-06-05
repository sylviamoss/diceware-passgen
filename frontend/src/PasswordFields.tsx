import React from 'react';
import './PasswordFields.css';

export interface Pass {
    words: string,
    password: string
}

type PasswordFieldsProps = {
    isOpen: boolean,
    pass: Pass
}

const PasswordFields: React.FC<PasswordFieldsProps> = ({ isOpen, pass }) => {
    if (isOpen) {
        return (
            <div className="password-fields">
                <div className="field" >
                    <input className="form-control" value={pass.words} readOnly></input>
                    <button type="button" className="btn btn-outline-warning">copy</button>
                </div>
                <div className="field" >
                    <input className="form-control" value={pass.password} readOnly></input>
                    <button type="button" className="btn btn-outline-warning">copy</button>
                </div>
            </div>
        )
    }
    return null
}

export default PasswordFields;
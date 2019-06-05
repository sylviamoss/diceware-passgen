import React from 'react';
import './PasswordFields.css';

export interface Words {
    plain: string,
    password: string
}

type PasswordFieldsProps = {
    isOpen: boolean,
    words: Words
}

const PasswordFields: React.FC<PasswordFieldsProps> = ({ isOpen, words }) => {
    if (isOpen) {
        return (
            <div className="password-fields">
                <div className="field" >
                    <input className="form-control" value={words.plain} readOnly></input>
                    <button type="button" className="btn btn-outline-warning">copy</button>
                </div>
                <div className="field" >
                    <input className="form-control" value={words.password} readOnly></input>
                    <button type="button" className="btn btn-outline-warning">copy</button>
                </div>
            </div>
        )
    }
    return null
}

export default PasswordFields;
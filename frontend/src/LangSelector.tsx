import React from 'react';
import './LangSelector.css';

type LangSelectorProps = {
    handleLangSelector: (lang: string) => void
}

const LangSelector: React.FC<LangSelectorProps> = ({handleLangSelector}) => {
    return (
        <div className="lang-selectors">
            <button type="button" className="btn btn-outline-info" onClick={() => handleLangSelector("pt")}>Portuguese</button>
            <button type="button" className="btn btn-outline-info" onClick={() => handleLangSelector("en")}>English</button>
        </div>
    )
}

export default LangSelector;

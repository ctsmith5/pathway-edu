import React, { useState } from 'react';
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import { vscDarkPlus } from 'react-syntax-highlighter/dist/esm/styles/prism';

const languageLabels = {
  csharp: 'C#',
  typescript: 'TypeScript',
  python: 'Python',
};

const languageDisplayOrder = ['csharp', 'typescript', 'python'];

const MultiCodeBlock = ({ data }) => {
  const { codes } = data;
  const [selectedLanguage, setSelectedLanguage] = useState('csharp');

  if (!codes || typeof codes !== 'object') {
    return null;
  }

  // Filter to only show languages that exist in the data
  const availableLanguages = languageDisplayOrder.filter(lang => codes[lang]);

  if (availableLanguages.length === 0) {
    return null;
  }

  // If only one language, just show it without toggle
  if (availableLanguages.length === 1) {
    const lang = availableLanguages[0];
    return (
      <div className="code-block">
        <SyntaxHighlighter
          language={lang === 'csharp' ? 'csharp' : lang}
          style={vscDarkPlus}
          className="code-block__highlighter"
        >
          {codes[lang]}
        </SyntaxHighlighter>
      </div>
    );
  }

  return (
    <div className="multi-code-block">
      <div className="multi-code-block__tabs">
        {availableLanguages.map((lang) => (
          <button
            key={lang}
            className={`multi-code-block__tab ${selectedLanguage === lang ? 'multi-code-block__tab--active' : ''}`}
            onClick={() => setSelectedLanguage(lang)}
          >
            {languageLabels[lang] || lang}
          </button>
        ))}
      </div>
      <div className="multi-code-block__content">
        <SyntaxHighlighter
          language={selectedLanguage === 'csharp' ? 'csharp' : selectedLanguage}
          style={vscDarkPlus}
          className="code-block__highlighter"
        >
          {codes[selectedLanguage]}
        </SyntaxHighlighter>
      </div>
    </div>
  );
};

export default MultiCodeBlock;
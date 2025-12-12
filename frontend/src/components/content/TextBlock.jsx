import React from 'react';
import ReactMarkdown from 'react-markdown';

const TextBlock = ({ data }) => {
  if (!data?.markdown) {
    return null;
  }

  return (
    <div className="text-block">
      <ReactMarkdown>{data.markdown}</ReactMarkdown>
    </div>
  );
};

export default TextBlock;


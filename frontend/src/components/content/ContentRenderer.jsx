import React from 'react';
import TextBlock from './TextBlock';
import CodeBlock from './CodeBlock';
import ImageBlock from './ImageBlock';
import CalloutBlock from './CalloutBlock';
import ExerciseBlock from './ExerciseBlock';
import VideoBlock from './VideoBlock';
import './content.css';

const blockComponents = {
  text: TextBlock,
  code: CodeBlock,
  image: ImageBlock,
  callout: CalloutBlock,
  exercise: ExerciseBlock,
  video: VideoBlock,
};

const ContentRenderer = ({ content }) => {
  if (!content || !Array.isArray(content)) {
    return null;
  }

  return (
    <div className="content-renderer">
      {content.map((block, index) => {
        const BlockComponent = blockComponents[block.type];
        
        if (!BlockComponent) {
          console.warn(`Unknown block type: ${block.type}`);
          return null;
        }

        return (
          <div key={index} className={`content-block content-block--${block.type}`}>
            <BlockComponent data={block.data} />
          </div>
        );
      })}
    </div>
  );
};

export default ContentRenderer;


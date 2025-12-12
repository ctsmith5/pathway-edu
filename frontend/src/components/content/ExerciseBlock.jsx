import React, { useState } from 'react';

const ExerciseBlock = ({ data }) => {
  const [showSolution, setShowSolution] = useState(false);
  const [showHints, setShowHints] = useState(false);

  if (!data?.prompt) {
    return null;
  }

  const { prompt, solution, hints = [] } = data;

  return (
    <div className="exercise-block">
      <div className="exercise-block__header">
        <span className="exercise-block__icon">✏️</span>
        <span className="exercise-block__title">Exercise</span>
      </div>
      
      <div className="exercise-block__prompt">
        <p>{prompt}</p>
      </div>

      <div className="exercise-block__actions">
        {hints.length > 0 && (
          <button 
            className="exercise-block__btn exercise-block__btn--hint"
            onClick={() => setShowHints(!showHints)}
          >
            {showHints ? 'Hide Hints' : 'Show Hints'}
          </button>
        )}
        
        {solution && (
          <button 
            className="exercise-block__btn exercise-block__btn--solution"
            onClick={() => setShowSolution(!showSolution)}
          >
            {showSolution ? 'Hide Solution' : 'Reveal Solution'}
          </button>
        )}
      </div>

      {showHints && hints.length > 0 && (
        <div className="exercise-block__hints">
          <strong>Hints:</strong>
          <ul>
            {hints.map((hint, index) => (
              <li key={index}>{hint}</li>
            ))}
          </ul>
        </div>
      )}

      {showSolution && solution && (
        <div className="exercise-block__solution">
          <strong>Solution:</strong>
          <pre>{solution}</pre>
        </div>
      )}
    </div>
  );
};

export default ExerciseBlock;


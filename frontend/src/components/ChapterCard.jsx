import React from 'react';
import { Link } from 'react-router-dom';
import './ChapterCard.css';

const ChapterCard = ({ chapter, completedModules = [], coursesMap = {} }) => {
  const { id, title, description, estimatedTime, items } = chapter;

  // Calculate progress based on completed modules
  const completedCount = items.filter(item => 
    completedModules.includes(item.moduleId)
  ).length;
  const progressPercent = items.length > 0 
    ? Math.round((completedCount / items.length) * 100) 
    : 0;
  const isCompleted = completedCount === items.length && items.length > 0;

  // Find the next uncompleted item
  const nextItem = items.find(item => !completedModules.includes(item.moduleId));

  // Get course ID for the next item
  const getModuleLink = (item) => {
    const course = Object.values(coursesMap).find(c => 
      c.modules?.some(m => m.id === item.moduleId)
    );
    if (course) {
      return `/module/${course.id}/${item.moduleId}`;
    }
    return null;
  };

  return (
    <div className={`chapter-card ${isCompleted ? 'chapter-card--completed' : ''}`}>
      <div className="chapter-card__header">
        <div className="chapter-card__title-row">
          <h3 className="chapter-card__title">{title}</h3>
          {isCompleted && <span className="chapter-card__badge">Complete</span>}
        </div>
        <span className="chapter-card__time">{estimatedTime}</span>
      </div>

      <p className="chapter-card__description">{description}</p>

      <div className="chapter-card__progress">
        <div className="progress-bar">
          <div 
            className="progress-bar__fill" 
            style={{ width: `${progressPercent}%` }}
          />
        </div>
        <span className="progress-text">
          {completedCount} of {items.length} complete ({progressPercent}%)
        </span>
      </div>

      <div className="chapter-card__actions">
        {nextItem ? (
          <>
            <span className="chapter-card__next-label">Next up:</span>
            <span className="chapter-card__next-title">{nextItem.title}</span>
          </>
        ) : (
          <span className="chapter-card__all-done">All items completed!</span>
        )}
      </div>

      <Link 
        to={`/chapter/${id}`} 
        className="chapter-card__button"
      >
        {isCompleted ? 'Review Chapter' : completedCount > 0 ? 'Continue' : 'Start Chapter'}
      </Link>
    </div>
  );
};

export default ChapterCard;


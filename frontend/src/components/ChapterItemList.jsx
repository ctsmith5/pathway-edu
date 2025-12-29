import React from 'react';
import { Link } from 'react-router-dom';
import './ChapterItemList.css';

const ChapterItemList = ({ items, completedModules = [], coursesMap = {} }) => {
  // Find the first uncompleted item index
  const firstUncompletedIndex = items.findIndex(
    item => !completedModules.includes(item.moduleId)
  );

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
    <div className="chapter-items">
      {items.map((item, index) => {
        const isCompleted = completedModules.includes(item.moduleId);
        const isLocked = !isCompleted && index > firstUncompletedIndex && firstUncompletedIndex !== -1;
        const isCurrent = index === firstUncompletedIndex;
        const moduleLink = getModuleLink(item);

        return (
          <div 
            key={item.moduleId}
            className={`chapter-item ${isCompleted ? 'chapter-item--completed' : ''} ${isLocked ? 'chapter-item--locked' : ''} ${isCurrent ? 'chapter-item--current' : ''}`}
          >
            <div className="chapter-item__number">
              {isCompleted ? (
                <span className="chapter-item__check">✓</span>
              ) : isLocked ? (
                <span className="chapter-item__lock">🔒</span>
              ) : (
                <span>{index + 1}</span>
              )}
            </div>
            
            <div className="chapter-item__content">
              <div className="chapter-item__header">
                <h3 className="chapter-item__title">{item.title}</h3>
                <span className="chapter-item__course">{item.courseTitle}</span>
              </div>
              {item.description && (
                <p className="chapter-item__description">{item.description}</p>
              )}
            </div>

            <div className="chapter-item__action">
              {isLocked ? (
                <span className="chapter-item__locked-text">Complete previous items first</span>
              ) : moduleLink ? (
                <Link 
                  to={moduleLink}
                  className={`chapter-item__button ${isCompleted ? 'chapter-item__button--review' : ''}`}
                >
                  {isCompleted ? 'Review' : isCurrent ? 'Start' : 'View'}
                </Link>
              ) : (
                <span className="chapter-item__coming-soon">Coming soon</span>
              )}
            </div>
          </div>
        );
      })}
    </div>
  );
};

export default ChapterItemList;


import React from 'react';
import { Link } from 'react-router-dom';
import './CourseProgressCard.css';

const CourseProgressCard = ({ courseWithProgress }) => {
  const { course, completed_modules = [], progress_percent = 0, is_completed } = courseWithProgress;
  const { id, title, description, modules = [] } = course;

  // Find next uncompleted modules
  const uncompletedModules = modules.filter(
    (module) => !completed_modules.includes(module.id)
  );
  const nextModules = uncompletedModules.slice(0, 3);

  // Determine course status
  const hasModules = modules.length > 0;
  const isStarted = completed_modules.length > 0;

  return (
    <div className={`course-card ${is_completed ? 'course-card--completed' : ''}`}>
      <div className="course-card__header">
        <h3 className="course-card__title">{title}</h3>
        {is_completed && <span className="course-card__badge">Completed</span>}
      </div>

      <p className="course-card__description">{description}</p>

      {hasModules && (
        <>
          <div className="course-card__progress">
            <div className="progress-bar">
              <div 
                className="progress-bar__fill" 
                style={{ width: `${progress_percent}%` }}
              />
            </div>
            <span className="progress-text">
              {Math.round(progress_percent)}% complete
            </span>
          </div>

          <div className="course-card__modules">
            <h4>{isStarted ? 'Next Up' : 'Get Started'}</h4>
            {nextModules.length > 0 ? (
              <ul className="module-list">
                {nextModules.map((module, index) => (
                  <li key={module.id}>
                    <Link 
                      to={`/module/${id}/${module.id}`}
                      className="module-link"
                    >
                      <span className="module-number">{index + 1}</span>
                      <span className="module-title">{module.title}</span>
                    </Link>
                  </li>
                ))}
              </ul>
            ) : (
              <p className="all-complete">All modules completed!</p>
            )}
          </div>
        </>
      )}

      {!hasModules && (
        <div className="course-card__empty">
          <p>Modules coming soon...</p>
        </div>
      )}
    </div>
  );
};

export default CourseProgressCard;


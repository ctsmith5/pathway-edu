import React from 'react';
import './Badge.css';

const Badge = ({ course, isEarned = false }) => {
  const { title, description } = course;

  // Generate a consistent color based on course title
  const getColor = (str) => {
    const colors = [
      { bg: '#6366f1', light: '#e0e7ff' }, // Indigo
      { bg: '#8b5cf6', light: '#ede9fe' }, // Violet
      { bg: '#ec4899', light: '#fce7f3' }, // Pink
      { bg: '#f97316', light: '#ffedd5' }, // Orange
      { bg: '#14b8a6', light: '#ccfbf1' }, // Teal
      { bg: '#3b82f6', light: '#dbeafe' }, // Blue
      { bg: '#22c55e', light: '#dcfce7' }, // Green
    ];
    const index = str.split('').reduce((acc, char) => acc + char.charCodeAt(0), 0) % colors.length;
    return colors[index];
  };

  const color = getColor(title);

  // Get initials for the badge icon
  const getInitials = (str) => {
    return str.split(' ')
      .map(word => word[0])
      .join('')
      .toUpperCase()
      .slice(0, 2);
  };

  return (
    <div className={`badge ${isEarned ? 'badge--earned' : 'badge--locked'}`}>
      <div 
        className="badge__icon"
        style={{ 
          backgroundColor: isEarned ? color.bg : '#9ca3af',
          color: 'white'
        }}
      >
        {isEarned ? (
          <span className="badge__initials">{getInitials(title)}</span>
        ) : (
          <span className="badge__lock">🔒</span>
        )}
      </div>
      <div className="badge__info">
        <h3 className="badge__title">{title}</h3>
        <p className="badge__description">
          {isEarned ? 'Course completed!' : description || 'Complete all modules to earn this badge'}
        </p>
      </div>
      {isEarned && (
        <div className="badge__earned-indicator">
          <span>✓</span>
        </div>
      )}
    </div>
  );
};

export default Badge;


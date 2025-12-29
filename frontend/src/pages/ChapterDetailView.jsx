import React, { useState, useEffect } from 'react';
import { useParams, Link } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import { API_URL } from '../config/api';
import { getChapterById } from '../data/chapters';
import ChapterItemList from '../components/ChapterItemList';
import './ChapterDetailView.css';

const ChapterDetailView = () => {
  const { chapterId } = useParams();
  const { token } = useAuth();
  const [coursesWithProgress, setCoursesWithProgress] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const chapter = getChapterById(chapterId);

  useEffect(() => {
    const fetchProgress = async () => {
      try {
        const response = await fetch(`${API_URL}/user/progress`, {
          headers: {
            'Authorization': `Bearer ${token}`
          }
        });

        if (!response.ok) {
          throw new Error('Failed to fetch progress');
        }

        const data = await response.json();
        setCoursesWithProgress(data || []);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    if (token) {
      fetchProgress();
    }
  }, [token]);

  // Create a map of courses by title for easy lookup
  const coursesMap = {};
  coursesWithProgress.forEach(cwp => {
    coursesMap[cwp.course.title] = cwp.course;
  });

  // Get all completed modules across all courses
  const allCompletedModules = coursesWithProgress.flatMap(
    cwp => cwp.completed_modules || []
  );

  if (!chapter) {
    return (
      <div className="chapter-detail">
        <div className="chapter-detail__error">
          <h2>Chapter Not Found</h2>
          <p>The chapter you're looking for doesn't exist.</p>
          <Link to="/dashboard" className="back-link">Back to Dashboard</Link>
        </div>
      </div>
    );
  }

  // Calculate progress
  const completedCount = chapter.items.filter(
    item => allCompletedModules.includes(item.moduleId)
  ).length;
  const progressPercent = chapter.items.length > 0
    ? Math.round((completedCount / chapter.items.length) * 100)
    : 0;
  const isCompleted = completedCount === chapter.items.length && chapter.items.length > 0;

  if (loading) {
    return (
      <div className="chapter-detail">
        <div className="chapter-detail__loading">
          <div className="loading-spinner"></div>
          <p>Loading chapter...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="chapter-detail">
        <div className="chapter-detail__error">
          <p>Error: {error}</p>
          <button onClick={() => window.location.reload()}>Try Again</button>
        </div>
      </div>
    );
  }

  return (
    <div className="chapter-detail">
      <div className="chapter-detail__container">
        {/* Breadcrumb */}
        <nav className="chapter-detail__breadcrumb">
          <Link to="/dashboard">Dashboard</Link>
          <span className="breadcrumb-separator">/</span>
          <span className="breadcrumb-current">{chapter.title}</span>
        </nav>

        {/* Header */}
        <header className="chapter-detail__header">
          <div className="chapter-detail__title-row">
            <h1>{chapter.title}</h1>
            {isCompleted && <span className="chapter-detail__badge">Complete</span>}
          </div>
          <p className="chapter-detail__description">{chapter.description}</p>
          
          <div className="chapter-detail__meta">
            <span className="chapter-detail__time">
              <span className="meta-icon">⏱</span>
              {chapter.estimatedTime}
            </span>
            <span className="chapter-detail__count">
              <span className="meta-icon">📚</span>
              {chapter.items.length} items
            </span>
          </div>

          <div className="chapter-detail__progress">
            <div className="progress-bar">
              <div 
                className={`progress-bar__fill ${isCompleted ? 'progress-bar__fill--complete' : ''}`}
                style={{ width: `${progressPercent}%` }}
              />
            </div>
            <span className="progress-text">
              {completedCount} of {chapter.items.length} complete ({progressPercent}%)
            </span>
          </div>
        </header>

        {/* Items List */}
        <section className="chapter-detail__content">
          <h2>Chapter Contents</h2>
          <ChapterItemList 
            items={chapter.items}
            completedModules={allCompletedModules}
            coursesMap={coursesMap}
          />
        </section>

        {/* Back to Dashboard */}
        <footer className="chapter-detail__footer">
          <Link to="/dashboard" className="chapter-detail__back">
            ← Back to Dashboard
          </Link>
        </footer>
      </div>
    </div>
  );
};

export default ChapterDetailView;


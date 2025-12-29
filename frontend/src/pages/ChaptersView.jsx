import React, { useState, useEffect } from 'react';
import { useAuth } from '../context/AuthContext';
import { API_URL } from '../config/api';
import { chapters } from '../data/chapters';
import ChapterCard from '../components/ChapterCard';
import './ChaptersView.css';

const ChaptersView = () => {
  const { user, token } = useAuth();
  const [coursesWithProgress, setCoursesWithProgress] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

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

  // Calculate total stats
  const totalChapters = chapters.length;
  const completedChapters = chapters.filter(chapter => 
    chapter.items.every(item => allCompletedModules.includes(item.moduleId))
  ).length;
  const totalItems = chapters.reduce((acc, ch) => acc + ch.items.length, 0);
  const completedItems = chapters.reduce((acc, ch) => 
    acc + ch.items.filter(item => allCompletedModules.includes(item.moduleId)).length, 0
  );

  if (loading) {
    return (
      <div className="chapters-view">
        <div className="chapters-loading">
          <div className="loading-spinner"></div>
          <p>Loading your progress...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="chapters-view">
        <div className="chapters-error">
          <p>Error: {error}</p>
          <button onClick={() => window.location.reload()}>Try Again</button>
        </div>
      </div>
    );
  }

  return (
    <div className="chapters-view">
      <div className="chapters-container">
        <header className="chapters-header">
          <h1>Welcome back, {user?.name?.split(' ')[0]}!</h1>
          <p>Continue your learning journey through the chapters below.</p>
        </header>

        <section className="chapters-stats">
          <div className="stat-card">
            <span className="stat-value">{completedChapters}/{totalChapters}</span>
            <span className="stat-label">Chapters</span>
          </div>
          <div className="stat-card">
            <span className="stat-value">{completedItems}/{totalItems}</span>
            <span className="stat-label">Items Completed</span>
          </div>
          <div className="stat-card">
            <span className="stat-value">
              {totalItems > 0 ? Math.round((completedItems / totalItems) * 100) : 0}%
            </span>
            <span className="stat-label">Overall Progress</span>
          </div>
        </section>

        <section className="chapters-list">
          <h2>Your Learning Path</h2>
          <div className="chapters-grid">
            {chapters.map((chapter) => (
              <ChapterCard
                key={chapter.id}
                chapter={chapter}
                completedModules={allCompletedModules}
                coursesMap={coursesMap}
              />
            ))}
          </div>
        </section>
      </div>
    </div>
  );
};

export default ChaptersView;


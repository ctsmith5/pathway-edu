import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import { API_URL } from '../config/api';
import Badge from '../components/Badge';
import './ProfileView.css';

const ProfileView = () => {
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

  // Calculate stats
  const totalCourses = coursesWithProgress.length;
  const completedCourses = coursesWithProgress.filter(c => c.is_completed).length;
  const completedModules = coursesWithProgress.reduce(
    (acc, c) => acc + (c.completed_modules?.length || 0), 0
  );

  if (loading) {
    return (
      <div className="profile-view">
        <div className="profile-loading">
          <div className="loading-spinner"></div>
          <p>Loading your profile...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="profile-view">
        <div className="profile-error">
          <p>Error: {error}</p>
          <button onClick={() => window.location.reload()}>Try Again</button>
        </div>
      </div>
    );
  }

  return (
    <div className="profile-view">
      <div className="profile-container">
        {/* Header */}
        <header className="profile-header">
          <div className="profile-avatar">
            {user?.name?.charAt(0).toUpperCase() || '?'}
          </div>
          <div className="profile-info">
            <h1>{user?.name || 'Student'}</h1>
            <p className="profile-email">{user?.email}</p>
          </div>
        </header>

        {/* Stats */}
        <section className="profile-stats">
          <div className="stat-card">
            <span className="stat-value">{completedCourses}</span>
            <span className="stat-label">Courses Completed</span>
          </div>
          <div className="stat-card">
            <span className="stat-value">{completedModules}</span>
            <span className="stat-label">Modules Completed</span>
          </div>
          <div className="stat-card">
            <span className="stat-value">
              {totalCourses > 0 ? Math.round((completedCourses / totalCourses) * 100) : 0}%
            </span>
            <span className="stat-label">Overall Progress</span>
          </div>
        </section>

        {/* Badges Section */}
        <section className="profile-badges">
          <h2>Your Badges</h2>
          <p className="badges-subtitle">
            Complete all modules in a course to earn its badge
          </p>
          
          <div className="badges-grid">
            {coursesWithProgress.map((cwp) => (
              <Badge
                key={cwp.course.id}
                course={cwp.course}
                isEarned={cwp.is_completed}
              />
            ))}
          </div>

          {coursesWithProgress.length === 0 && (
            <div className="badges-empty">
              <p>No courses available yet.</p>
            </div>
          )}
        </section>

        {/* Course Progress Section */}
        <section className="profile-progress">
          <h2>Course Progress</h2>
          <div className="progress-list">
            {coursesWithProgress.map((cwp) => {
              const totalMods = cwp.course.modules?.length || 0;
              const completedMods = cwp.completed_modules?.length || 0;
              const percent = totalMods > 0 ? Math.round((completedMods / totalMods) * 100) : 0;

              return (
                <div key={cwp.course.id} className="progress-item">
                  <div className="progress-item__header">
                    <span className="progress-item__title">{cwp.course.title}</span>
                    <span className="progress-item__count">
                      {completedMods}/{totalMods} modules
                    </span>
                  </div>
                  <div className="progress-item__bar">
                    <div 
                      className={`progress-item__fill ${cwp.is_completed ? 'progress-item__fill--complete' : ''}`}
                      style={{ width: `${percent}%` }}
                    />
                  </div>
                </div>
              );
            })}
          </div>
        </section>

        {/* Back Link */}
        <footer className="profile-footer">
          <Link to="/dashboard" className="profile-back">
            ← Back to Dashboard
          </Link>
        </footer>
      </div>
    </div>
  );
};

export default ProfileView;


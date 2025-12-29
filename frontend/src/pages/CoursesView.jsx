import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import { API_URL } from '../config/api';
import CourseProgressCard from '../components/CourseProgressCard';
import './CoursesView.css';

const CoursesView = () => {
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

  if (loading) {
    return (
      <div className="courses-view">
        <div className="courses-loading">
          <div className="loading-spinner"></div>
          <p>Loading courses...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="courses-view">
        <div className="courses-error">
          <p>Error: {error}</p>
          <button onClick={() => window.location.reload()}>Try Again</button>
        </div>
      </div>
    );
  }

  return (
    <div className="courses-view">
      <div className="courses-container">
        {/* Breadcrumb */}
        <nav className="courses-breadcrumb">
          <Link to="/dashboard">Dashboard</Link>
          <span className="breadcrumb-separator">/</span>
          <span className="breadcrumb-current">All Courses</span>
        </nav>

        <header className="courses-header">
          <h1>All Courses</h1>
          <p>Browse all available courses and track your progress across each one.</p>
        </header>

        <section className="courses-stats">
          <div className="stat-card">
            <span className="stat-value">{coursesWithProgress.length}</span>
            <span className="stat-label">Total Courses</span>
          </div>
          <div className="stat-card">
            <span className="stat-value">
              {coursesWithProgress.reduce((acc, c) => acc + (c.completed_modules?.length || 0), 0)}
            </span>
            <span className="stat-label">Modules Completed</span>
          </div>
          <div className="stat-card">
            <span className="stat-value">
              {coursesWithProgress.filter(c => c.is_completed).length}
            </span>
            <span className="stat-label">Courses Completed</span>
          </div>
        </section>

        <section className="courses-list">
          <div className="courses-grid">
            {coursesWithProgress.map((cwp) => (
              <CourseProgressCard 
                key={cwp.course.id} 
                courseWithProgress={cwp}
              />
            ))}
          </div>
        </section>

        <footer className="courses-footer">
          <Link to="/dashboard" className="courses-back">
            ← Back to Dashboard
          </Link>
        </footer>
      </div>
    </div>
  );
};

export default CoursesView;


import React, { useState, useEffect } from 'react';
import { useAuth } from '../context/AuthContext';
import { API_URL } from '../config/api';
import CourseProgressCard from '../components/CourseProgressCard';
import './Dashboard.css';

const Dashboard = () => {
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
      <div className="dashboard">
        <div className="dashboard-loading">
          <div className="loading-spinner"></div>
          <p>Loading your courses...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="dashboard">
        <div className="dashboard-error">
          <p>Error: {error}</p>
          <button onClick={() => window.location.reload()}>Try Again</button>
        </div>
      </div>
    );
  }

  return (
    <div className="dashboard">
      <div className="dashboard-container">
        <header className="dashboard-header">
          <h1>Welcome back, {user?.name?.split(' ')[0]}!</h1>
          <p>Continue your learning journey. Pick up where you left off.</p>
        </header>

        <section className="dashboard-stats">
          <div className="stat-card">
            <span className="stat-value">{coursesWithProgress.length}</span>
            <span className="stat-label">Courses</span>
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

        <section className="dashboard-courses">
          <h2>Your Courses</h2>
          <div className="courses-grid">
            {coursesWithProgress.map((cwp) => (
              <CourseProgressCard 
                key={cwp.course.id} 
                courseWithProgress={cwp}
              />
            ))}
          </div>
        </section>
      </div>
    </div>
  );
};

export default Dashboard;


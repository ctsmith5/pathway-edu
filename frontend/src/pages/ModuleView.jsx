import React, { useState, useEffect } from 'react';
import { useParams, useNavigate, Link } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import { API_URL } from '../config/api';
import { ContentRenderer } from '../components/content';
import './ModuleView.css';

const ModuleView = () => {
  const { courseId, moduleId } = useParams();
  const navigate = useNavigate();
  const { token } = useAuth();
  
  const [course, setCourse] = useState(null);
  const [currentModule, setCurrentModule] = useState(null);
  const [moduleIndex, setModuleIndex] = useState(-1);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [isCompleted, setIsCompleted] = useState(false);
  const [completing, setCompleting] = useState(false);

  useEffect(() => {
    const fetchCourseAndProgress = async () => {
      try {
        // Fetch course
        const courseResponse = await fetch(`${API_URL}/courses/${courseId}`);
        
        if (!courseResponse.ok) {
          throw new Error('Course not found');
        }

        const courseData = await courseResponse.json();
        setCourse(courseData);

        // Find the current module
        const index = courseData.modules.findIndex(m => m.id === moduleId);
        if (index === -1) {
          throw new Error('Module not found');
        }

        setModuleIndex(index);
        setCurrentModule(courseData.modules[index]);

        // Fetch user progress to check if module is completed
        if (token) {
          const progressResponse = await fetch(`${API_URL}/user/progress`, {
            headers: {
              'Authorization': `Bearer ${token}`
            }
          });
          if (progressResponse.ok) {
            const progressData = await progressResponse.json();
            const courseProgress = progressData.find(p => p.course.id === courseId);
            if (courseProgress && courseProgress.completed_modules?.includes(moduleId)) {
              setIsCompleted(true);
            }
          }
        }
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchCourseAndProgress();
  }, [courseId, moduleId, token]);

  const handleMarkComplete = async () => {
    if (completing || isCompleted) return;
    
    setCompleting(true);
    try {
      const response = await fetch(`${API_URL}/user/progress/complete`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify({
          course_id: courseId,
          module_id: moduleId
        })
      });

      if (response.ok) {
        setIsCompleted(true);
      } else {
        console.error('Failed to mark module as complete');
      }
    } catch (err) {
      console.error('Error marking module complete:', err);
    } finally {
      setCompleting(false);
    }
  };

  const handlePrevious = () => {
    if (moduleIndex > 0 && course) {
      const prevModule = course.modules[moduleIndex - 1];
      navigate(`/module/${courseId}/${prevModule.id}`);
    }
  };

  const handleNext = () => {
    if (course && moduleIndex < course.modules.length - 1) {
      const nextModule = course.modules[moduleIndex + 1];
      navigate(`/module/${courseId}/${nextModule.id}`);
    }
  };

  if (loading) {
    return (
      <div className="module-view">
        <div className="module-loading">
          <div className="loading-spinner"></div>
          <p>Loading module...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="module-view">
        <div className="module-error">
          <h2>Oops!</h2>
          <p>{error}</p>
          <Link to="/dashboard" className="back-link">Back to Dashboard</Link>
        </div>
      </div>
    );
  }

  const hasPrevious = moduleIndex > 0;
  const hasNext = course && moduleIndex < course.modules.length - 1;

  return (
    <div className="module-view">
      <div className="module-view__container">
        {/* Breadcrumb Navigation */}
        <nav className="module-breadcrumb">
          <Link to="/dashboard">Dashboard</Link>
          <span className="breadcrumb-separator">/</span>
          <span className="breadcrumb-course">{course?.title}</span>
          <span className="breadcrumb-separator">/</span>
          <span className="breadcrumb-current">{currentModule?.title}</span>
        </nav>

        {/* Module Header */}
        <header className="module-header">
          <div className="module-header__info">
            <span className="module-number">
              Module {moduleIndex + 1} of {course?.modules.length}
            </span>
            <h1>{currentModule?.title}</h1>
          </div>
        </header>

        {/* Module Content */}
        <main className="module-content">
          {currentModule?.content && currentModule.content.length > 0 ? (
            <ContentRenderer content={currentModule.content} />
          ) : (
            <div className="module-empty">
              <p>Content coming soon...</p>
            </div>
          )}

          {/* Complete Module Button */}
          <div className="module-complete-section">
            {isCompleted ? (
              <div className="module-complete-badge">
                <span className="complete-icon">✓</span>
                <span>Module Completed!</span>
              </div>
            ) : (
              <button 
                className="module-complete-btn"
                onClick={handleMarkComplete}
                disabled={completing}
              >
                {completing ? 'Marking Complete...' : 'Mark as Complete'}
              </button>
            )}
          </div>
        </main>

        {/* Navigation Footer */}
        <footer className="module-navigation">
          <button 
            className="nav-btn nav-btn--prev"
            onClick={handlePrevious}
            disabled={!hasPrevious}
          >
            <span className="nav-arrow">←</span>
            <span className="nav-label">
              {hasPrevious ? course.modules[moduleIndex - 1].title : 'Previous'}
            </span>
          </button>

          <Link to="/dashboard" className="nav-btn nav-btn--dashboard">
            Back to Dashboard
          </Link>

          <button 
            className="nav-btn nav-btn--next"
            onClick={handleNext}
            disabled={!hasNext}
          >
            <span className="nav-label">
              {hasNext ? course.modules[moduleIndex + 1].title : 'Next'}
            </span>
            <span className="nav-arrow">→</span>
          </button>
        </footer>
      </div>
    </div>
  );
};

export default ModuleView;


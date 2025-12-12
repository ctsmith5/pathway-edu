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

  useEffect(() => {
    const fetchCourse = async () => {
      try {
        const response = await fetch(`${API_URL}/courses/${courseId}`);
        
        if (!response.ok) {
          throw new Error('Course not found');
        }

        const courseData = await response.json();
        setCourse(courseData);

        // Find the current module
        const index = courseData.modules.findIndex(m => m.id === moduleId);
        if (index === -1) {
          throw new Error('Module not found');
        }

        setModuleIndex(index);
        setCurrentModule(courseData.modules[index]);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchCourse();
  }, [courseId, moduleId]);

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


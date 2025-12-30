import React from 'react';
import { Link } from 'react-router-dom';
import './LandingPage.css';

const LandingPage = () => {
  return (
    <div className="landing-page">
      {/* Hero Section */}
      <div className="hero-section">
        <div className="hero-overlay"></div>
        <div className="hero-content">
          <h1>Launch Your Child's Career in Tech</h1>
          <p className="hero-subtitle">
            Practical, job-ready software engineering skills for the next generation of innovators.
            Give them the tools to build the future.
          </p>
          <div className="hero-buttons">
            <Link to="/dashboard" className="cta-button primary">Get Started Today</Link>
            <Link to="/curriculum" className="cta-button secondary">View Curriculum</Link>
          </div>
        </div>
      </div>

      {/* Features Section */}
      <div className="features-section">
        <div className="section-header">
          <h2>Why Choose Pathway?</h2>
          <p>We bridge the gap between hobbyist coding and professional engineering.</p>
        </div>
        
        <div className="features-grid">
          <div className="feature-card">
            <div className="card-image-wrapper">
               <img src="https://images.unsplash.com/photo-1498050108023-c5249f4df085?auto=format&fit=crop&w=600&q=80" alt="Coding on laptop" />
            </div>
            <h3>Real-World Skills</h3>
            <p>
              We move beyond theory. Our curriculum is built around the actual tools and languages (Go, React, SQL) used by professional software engineers every day.
            </p>
          </div>
          
          <div className="feature-card">
            <div className="card-image-wrapper">
              <img src="https://images.unsplash.com/photo-1516321318423-f06f85e504b3?auto=format&fit=crop&w=600&q=80" alt="Mentoring" />
            </div>
            <h3>Guided Path</h3>
            <p>
              Many talented young adults just need direction. We provide a structured roadmap from basics to proficiency, eliminating the "tutorial hell" confusion.
            </p>
          </div>
          
          <div className="feature-card">
            <div className="card-image-wrapper">
              <img src="https://images.unsplash.com/photo-1551288049-bebda4e38f71?auto=format&fit=crop&w=600&q=80" alt="Data analytics" />
            </div>
            <h3>Track Progress</h3>
            <p>
              Monitor proficiency and milestones as they master the fundamentals of web development, backend systems, and databases.
            </p>
          </div>
        </div>
      </div>

      {/* About Section */}
      <div className="about-section">
        <div className="about-content">
          <div className="about-text">
            <h2>A Clear Alternative</h2>
            <p>
              The transition after high school can be uncertain. Pathway offers a clear alternative or supplement to traditional education, focusing strictly on employable technical skills.
            </p>
            <p>
              We empower students to build a portfolio that speaks for itself. No fluff, just code.
            </p>
            <Link to="/curriculum" className="cta-button primary">Learn More About Us</Link>
          </div>
          <div className="about-image">
             <img src="https://images.unsplash.com/photo-1522202176988-66273c2fd55f?auto=format&fit=crop&w=800&q=80" alt="Team working together" />
          </div>
        </div>
      </div>
    </div>
  );
};

export default LandingPage;

import React from 'react';
import { Link } from 'react-router-dom';
import './ViewCurriculum.css';

const ViewCurriculum = () => {
  const curriculumSections = [
    {
      title: "Getting Started",
      icon: "🚀",
      description: "Before you can build anything, you need the right tools and a place to save your work.",
      courses: [
        {
          name: "Development Tools",
          tagline: "Setting Up Your Workspace",
          why: "Every craftsperson needs their tools. We'll help you set up Visual Studio Code—a free, powerful text editor that professional developers use every day. Think of it as your workshop where all your projects will come to life.",
          outcome: "You'll have a professional coding environment ready to go, just like the pros use."
        },
        {
          name: "Git & GitHub",
          tagline: "Never Lose Your Work Again",
          why: "Imagine writing a 10-page essay and your computer crashes. Git is like an unlimited 'undo' button and automatic backup system for your code. GitHub is where developers store and share their work—it's like Google Docs, but for code. Every tech company uses it.",
          outcome: "You'll be able to save your progress, undo mistakes, and show your work to potential employers."
        }
      ]
    },
    {
      title: "Core Programming",
      icon: "💡",
      description: "Learn the fundamental building blocks that all software is built upon.",
      courses: [
        {
          name: "Code Concepts",
          tagline: "Speaking the Language of Computers",
          why: "Before you can write poetry, you need to learn the alphabet. We teach you how to give instructions to a computer in ways it understands—variables (storing information), functions (reusable instructions), and logic (making decisions).",
          outcome: "You'll understand how to think like a programmer and write code that actually does things."
        },
        {
          name: "HTTP & Networking",
          tagline: "How the Internet Actually Works",
          why: "Every time you load a website, your browser is having a conversation with a computer somewhere else. Understanding this conversation is essential for building anything on the web—from simple websites to complex apps.",
          outcome: "You'll understand what happens when you click a link, submit a form, or load an app."
        }
      ]
    },
    {
      title: "Building Things Right",
      icon: "🏗️",
      description: "Learn patterns and principles that separate amateur code from professional software.",
      courses: [
        {
          name: "Design Patterns",
          tagline: "Proven Solutions to Common Problems",
          why: "Developers have been solving the same problems for decades. Instead of reinventing the wheel, we teach you battle-tested approaches that make your code easier to understand, modify, and extend. It's like learning architectural blueprints for code.",
          outcome: "You'll write code that other developers can easily work with—a crucial skill for any job."
        },
        {
          name: "SOLID Principles",
          tagline: "Writing Code That Lasts",
          why: "Bad code works today but breaks tomorrow. SOLID principles are five simple rules that help you write software that's easy to change, fix, and grow. Companies value developers who write maintainable code because it saves them time and money.",
          outcome: "You'll build software that's flexible enough to evolve with changing requirements."
        }
      ]
    },
    {
      title: "Quality & Confidence",
      icon: "✅",
      description: "Learn how to prove your code works and catch bugs before users do.",
      courses: [
        {
          name: "Testing",
          tagline: "Making Sure Your Code Actually Works",
          why: "Professional developers don't just hope their code works—they prove it. We teach you how to write automated tests that check your code thousands of times faster than you could manually. This is what separates hobby projects from production-ready software.",
          outcome: "You'll ship code with confidence, knowing it won't break when users start using it."
        }
      ]
    },
    {
      title: "Working With Teams",
      icon: "👥",
      description: "Software is built by teams. Learn how professional teams collaborate and ship products.",
      courses: [
        {
          name: "SCRUM & Agile",
          tagline: "How Tech Teams Actually Work",
          why: "Almost every tech company uses 'Agile' methods to build software. SCRUM is the most popular approach—it's how teams break big projects into small pieces, stay organized, and deliver working software every few weeks instead of waiting months.",
          outcome: "You'll understand the rhythm of professional software development and be ready to join a team."
        }
      ]
    }
  ];

  return (
    <div className="curriculum-page">
      {/* Hero Section */}
      <div className="curriculum-hero">
        <div className="curriculum-hero__content">
          <h1>Our Curriculum</h1>
          <p className="curriculum-hero__subtitle">
            A practical roadmap from curious beginner to job-ready developer. 
            No fluff, no filler—just the skills that actually matter.
          </p>
        </div>
      </div>

      {/* Philosophy Section */}
      <section className="curriculum-philosophy">
        <div className="philosophy-content">
          <h2>Why These Skills?</h2>
          <p>
            We talked to hiring managers, reviewed job postings, and consulted with senior engineers 
            to build a curriculum focused on what actually gets you hired. Every topic we teach 
            is something you'll use on day one of a real job.
          </p>
          <div className="philosophy-points">
            <div className="philosophy-point">
              <span className="point-icon">🎯</span>
              <div>
                <h4>Practical First</h4>
                <p>We focus on doing, not just reading. You'll build real projects from day one.</p>
              </div>
            </div>
            <div className="philosophy-point">
              <span className="point-icon">📈</span>
              <div>
                <h4>Progressive Learning</h4>
                <p>Each skill builds on the last. No jumping into the deep end before you can swim.</p>
              </div>
            </div>
            <div className="philosophy-point">
              <span className="point-icon">💼</span>
              <div>
                <h4>Job-Ready</h4>
                <p>Everything we teach is used daily by professional developers at real companies.</p>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Curriculum Sections */}
      <section className="curriculum-sections">
        <div className="sections-container">
          {curriculumSections.map((section, sectionIndex) => (
            <div key={sectionIndex} className="curriculum-section">
              <div className="section-header">
                <span className="section-icon">{section.icon}</span>
                <div className="section-header-text">
                  <h2>{section.title}</h2>
                  <p>{section.description}</p>
                </div>
              </div>
              
              <div className="courses-list">
                {section.courses.map((course, courseIndex) => (
                  <div key={courseIndex} className="course-card">
                    <div className="course-card__header">
                      <h3>{course.name}</h3>
                      <span className="course-tagline">{course.tagline}</span>
                    </div>
                    <div className="course-card__body">
                      <div className="course-why">
                        <h4>Why Learn This?</h4>
                        <p>{course.why}</p>
                      </div>
                      <div className="course-outcome">
                        <h4>What You'll Achieve</h4>
                        <p>{course.outcome}</p>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            </div>
          ))}
        </div>
      </section>

      {/* CTA Section */}
      <section className="curriculum-cta">
        <div className="cta-content">
          <h2>Ready to Start Your Journey?</h2>
          <p>
            Join Pathway and get access to our full curriculum, hands-on projects, 
            and a clear path to becoming a professional developer.
          </p>
          <div className="cta-buttons">
            <Link to="/dashboard" className="cta-button primary">Get Started</Link>
            <Link to="/" className="cta-button secondary">Back to Home</Link>
          </div>
        </div>
      </section>
    </div>
  );
};

export default ViewCurriculum;

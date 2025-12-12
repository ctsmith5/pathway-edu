import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import './Navbar.css';

const CompassIcon = () => (
  <svg 
    width="38" 
    height="38" 
    viewBox="0 0 100 100" 
    fill="none" 
    xmlns="http://www.w3.org/2000/svg"
    className="compass-icon"
    style={{ transform: 'rotate(-20deg)' }}
  >
    {/* Drop shadow for depth */}
    <defs>
      <filter id="cartoonShadow" x="-20%" y="-20%" width="140%" height="140%">
        <feDropShadow dx="2" dy="3" stdDeviation="2" floodColor="#0f172a" floodOpacity="0.4"/>
      </filter>
      <linearGradient id="bodyGradient" x1="0%" y1="0%" x2="100%" y2="100%">
        <stop offset="0%" stopColor="#fbbf24" />
        <stop offset="100%" stopColor="#f59e0b" />
      </linearGradient>
      <linearGradient id="ringGradient" x1="0%" y1="0%" x2="100%" y2="100%">
        <stop offset="0%" stopColor="#854d0e" />
        <stop offset="100%" stopColor="#a16207" />
      </linearGradient>
      <linearGradient id="faceGradient" x1="0%" y1="0%" x2="0%" y2="100%">
        <stop offset="0%" stopColor="#fef3c7" />
        <stop offset="100%" stopColor="#fde68a" />
      </linearGradient>
      <linearGradient id="redNeedle" x1="50%" y1="0%" x2="50%" y2="100%">
        <stop offset="0%" stopColor="#ff6b6b" />
        <stop offset="100%" stopColor="#ee5253" />
      </linearGradient>
      <linearGradient id="blueNeedle" x1="50%" y1="0%" x2="50%" y2="100%">
        <stop offset="0%" stopColor="#74b9ff" />
        <stop offset="100%" stopColor="#0984e3" />
      </linearGradient>
    </defs>
    
    <g filter="url(#cartoonShadow)">
      {/* Outer gold ring - chunky cartoon style */}
      <circle cx="50" cy="50" r="44" fill="url(#ringGradient)" stroke="#713f12" strokeWidth="3" />
      
      {/* Inner compass face */}
      <circle cx="50" cy="50" r="36" fill="url(#faceGradient)" stroke="#ca8a04" strokeWidth="2" />
      
      {/* Decorative inner ring */}
      <circle cx="50" cy="50" r="30" fill="none" stroke="#d97706" strokeWidth="1.5" strokeDasharray="4 3" />
      
      {/* Cardinal direction dots - chunky */}
      <circle cx="50" cy="18" r="4" fill="#dc2626" stroke="#991b1b" strokeWidth="1.5" />
      <circle cx="50" cy="82" r="3" fill="#6b7280" stroke="#4b5563" strokeWidth="1" />
      <circle cx="18" cy="50" r="3" fill="#6b7280" stroke="#4b5563" strokeWidth="1" />
      <circle cx="82" cy="50" r="3" fill="#6b7280" stroke="#4b5563" strokeWidth="1" />
      
      {/* Compass needle - Red North */}
      <polygon 
        points="50,22 43,50 50,46 57,50" 
        fill="url(#redNeedle)"
        stroke="#b91c1c"
        strokeWidth="1.5"
        strokeLinejoin="round"
      />
      
      {/* Compass needle - Blue South */}
      <polygon 
        points="50,78 43,50 50,54 57,50" 
        fill="url(#blueNeedle)"
        stroke="#1d4ed8"
        strokeWidth="1.5"
        strokeLinejoin="round"
      />
      
      {/* Center pivot - chunky golden */}
      <circle cx="50" cy="50" r="7" fill="#fbbf24" stroke="#b45309" strokeWidth="2" />
      <circle cx="50" cy="50" r="3" fill="#fef3c7" />
      
      {/* Cartoon shine highlight */}
      <ellipse cx="35" cy="32" rx="8" ry="5" fill="white" opacity="0.5" transform="rotate(-30 35 32)" />
    </g>
  </svg>
);

const Navbar = ({ onLoginClick }) => {
  const { isAuthenticated, user, logout } = useAuth();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate('/');
  };

  return (
    <nav className="navbar">
      <div className="navbar-container">
        <Link to="/" className="navbar-logo">
          <CompassIcon />
          Pathway
        </Link>
        <ul className="nav-menu">
          {isAuthenticated ? (
            <>
              <li className="nav-item">
                <Link to="/dashboard" className="nav-links">
                  Dashboard
                </Link>
              </li>
              <li className="nav-item nav-user">
                <span className="nav-user-name">{user?.name}</span>
                <button 
                  className="nav-logout-btn"
                  onClick={handleLogout}
                >
                  Logout
                </button>
              </li>
            </>
          ) : (
            <li className="nav-item">
              <span className="nav-links" onClick={onLoginClick} role="button" tabIndex={0}>
                Login
              </span>
            </li>
          )}
        </ul>
      </div>
    </nav>
  );
};

export default Navbar;

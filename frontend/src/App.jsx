import React, { useState } from 'react'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Navbar from './components/Navbar'
import LandingPage from './pages/LandingPage'
import LoginModal from './components/LoginModal'
import './App.css'

function App() {
  const [isLoginOpen, setIsLoginOpen] = useState(false);

  const openLogin = () => setIsLoginOpen(true);
  const closeLogin = () => setIsLoginOpen(false);

  return (
    <Router>
      <div className="app">
        <Navbar onLoginClick={openLogin} />
        <LoginModal isOpen={isLoginOpen} onClose={closeLogin} />
        <Routes>
          <Route path="/" element={<LandingPage />} />
          {/* We can remove the /login route now since we have the modal, 
              or keep it as a fallback/admin route */}
        </Routes>
      </div>
    </Router>
  )
}

export default App

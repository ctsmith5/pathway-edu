import React, { useState } from 'react'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import { AuthProvider } from './context/AuthContext'
import Navbar from './components/Navbar'
import LandingPage from './pages/LandingPage'
import Dashboard from './pages/Dashboard'
import ModuleView from './pages/ModuleView'
import LoginModal from './components/LoginModal'
import ProtectedRoute from './components/ProtectedRoute'
import './App.css'

function App() {
  const [isLoginOpen, setIsLoginOpen] = useState(false);

  const openLogin = () => setIsLoginOpen(true);
  const closeLogin = () => setIsLoginOpen(false);

  return (
    <AuthProvider>
      <Router>
        <div className="app">
          <Navbar onLoginClick={openLogin} />
          <LoginModal isOpen={isLoginOpen} onClose={closeLogin} />
          <Routes>
            <Route path="/" element={<LandingPage />} />
            <Route 
              path="/dashboard" 
              element={
                <ProtectedRoute>
                  <Dashboard />
                </ProtectedRoute>
              } 
            />
            <Route 
              path="/module/:courseId/:moduleId" 
              element={
                <ProtectedRoute>
                  <ModuleView />
                </ProtectedRoute>
              } 
            />
          </Routes>
        </div>
      </Router>
    </AuthProvider>
  )
}

export default App

import React, { useState } from 'react'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import { AuthProvider } from './context/AuthContext'
import Navbar from './components/Navbar'
import LandingPage from './pages/LandingPage'
import ChaptersView from './pages/ChaptersView'
import ChapterDetailView from './pages/ChapterDetailView'
import CoursesView from './pages/CoursesView'
import ProfileView from './pages/ProfileView'
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
                  <ChaptersView />
                </ProtectedRoute>
              } 
            />
            <Route 
              path="/chapter/:chapterId" 
              element={
                <ProtectedRoute>
                  <ChapterDetailView />
                </ProtectedRoute>
              } 
            />
            <Route 
              path="/courses" 
              element={
                <ProtectedRoute>
                  <CoursesView />
                </ProtectedRoute>
              } 
            />
            <Route 
              path="/profile" 
              element={
                <ProtectedRoute>
                  <ProfileView />
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

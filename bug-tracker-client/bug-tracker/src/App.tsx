import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import NotFound404 from './components/NotFound404';
import Login from './components/Login';
import Dashboard from './components/Dashboard';
import Header from './components/Header';
import Footer from './components/Footer';
import './App.css';
import { useState } from 'react';
import { UserContext } from './context/userContext';

import DashIndex from './components/DashIndex';
import DashProjects from './components/DashProjects';
import DashTickets from './components/DashTickets';

function App() {
  const [user, setUser] = useState('');
  const value = { state: { user }, actions: { setUser } };
  return (
    <div className="App">
      <UserContext.Provider value={value}>
        <Header />
        <Router>
          <Routes>
            <Route path="/" element={<Login />} />
            <Route path="/dashboard" element={<Dashboard />}>
              {/*NESTED ROUTE DASHBOARD */}
              <Route index element={<DashIndex />} />
              <Route path="dashIndex" element={<DashIndex />} />
              <Route path="dashTickets" element={<DashTickets />} />
              <Route path="dashProjects" element={<DashProjects />} />
              {/*<Route path="*" element={<NoMatch />} />*/}
            </Route>
            <Route path="*" element={<NotFound404 />} />
          </Routes>
        </Router>

        <Footer />
      </UserContext.Provider>
    </div>
  );
}

export default App;

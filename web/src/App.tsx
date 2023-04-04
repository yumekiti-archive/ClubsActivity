import React, { VFC } from 'react';
import { AnimatePresence } from 'framer-motion';
import { Route, Routes, useLocation, useNavigate } from 'react-router-dom';
import Home from './components/pages/Home';
import Club from './components/pages/Club';

const App: VFC = () => {
  const location = useLocation();

  return (
    <AnimatePresence mode='wait'>
      <Routes location={location} key={location.pathname}>
        <Route path='/' element={<Home />} />
        <Route path='/clubs/:id' element={<Club />} />
      </Routes>
    </AnimatePresence>
  );
};

export default App;

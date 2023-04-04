import React, { VFC } from 'react';
import { AnimatePresence } from 'framer-motion';
import { Route, Routes, useLocation, useNavigate } from 'react-router-dom';
import { QueryClient, QueryClientProvider } from 'react-query';
import Home from './components/pages/Home';
import Club from './components/pages/Club';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
      retry: false,
    },
  },
});

const App: VFC = () => {
  const location = useLocation();

  return (
    <QueryClientProvider client={queryClient}>
      <AnimatePresence mode='wait'>
        <Routes location={location} key={location.pathname}>
          <Route path='/' element={<Home />} />
          <Route path='/clubs/:id' element={<Club />} />
        </Routes>
      </AnimatePresence>
    </QueryClientProvider>
  );
};

export default App;

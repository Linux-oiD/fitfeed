import React from 'react';
import '@ant-design/v5-patch-for-react-19';
import { ConfigProvider } from 'antd';
import { AuthProvider } from './context/AuthContext';
import MainLayout from './components/layout/MainLayout';
import Home from './pages/Home';
import './App.css';

const App: React.FC = () => {
  return (
    <ConfigProvider
      theme={{
        token: {
          colorPrimary: '#fc4c02', // FitFeed Orange (like Strava)
        },
      }}
    >
      <AuthProvider>
        <MainLayout>
          <Home />
        </MainLayout>
      </AuthProvider>
    </ConfigProvider>
  );
};

export default App;

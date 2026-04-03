import React from 'react';
import { Layout } from 'antd';

const { Footer } = Layout;

const AppFooter: React.FC = () => {
  return (
    <Footer style={{ textAlign: 'center', backgroundColor: '#f0f2f5', padding: '24px 50px' }}>
      <div style={{ marginBottom: '16px' }}>
        <strong>FITFEED</strong> - Privacy-First Fitness Social Platform
      </div>
      <div style={{ color: '#8c8c8c' }}>
        FitFeed ©{new Date().getFullYear()} Created with Ant Design
      </div>
    </Footer>
  );
};

export default AppFooter;

import React, { ReactNode } from 'react';
import { Layout } from 'antd';
import AppHeader from './AppHeader';
import AppFooter from './AppFooter';

const { Content } = Layout;

interface MainLayoutProps {
  children: ReactNode;
}

const MainLayout: React.FC<MainLayoutProps> = ({ children }) => {
  return (
    <Layout style={{ minHeight: '100vh', display: 'flex', flexDirection: 'column' }}>
      <AppHeader />
      <Content style={{ 
        flex: 1, 
        padding: '24px 50px', 
        maxWidth: '1200px', 
        margin: '0 auto', 
        width: '100%',
        backgroundColor: '#f0f2f5'
      }}>
        {children}
      </Content>
      <AppFooter />
    </Layout>
  );
};

export default MainLayout;

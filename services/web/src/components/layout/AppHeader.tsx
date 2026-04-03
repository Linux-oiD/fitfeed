import React, { useState } from 'react';
import { Layout, Menu, Button, Dropdown, Avatar, Space, Modal } from 'antd';
import { UserOutlined, SettingOutlined, LogoutOutlined, GoogleOutlined, IdcardOutlined } from '@ant-design/icons';
import { useAuth } from '../../context/AuthContext';

const { Header } = Layout;

const AppHeader: React.FC = () => {
  const { user, isLoggedIn, logout, config } = useAuth();
  const [isModalVisible, setIsModalVisible] = useState(false);

  const handleLogout = () => {
    logout();
  };

  const handlePasskeyLogin = () => {
    // TODO: Implement Passkey Login
    console.log("Passkey login initiated");
  };

  const handleOAuthLogin = (provider: string) => {
    if (!config) return;
    // Redirect to backend auth
    window.location.href = `${config.auth_url}/v1/oauth/${provider}/auth`;
  };

  const menuItems = [
    {
      key: 'user-info',
      disabled: true,
      label: (
        <Space>
          <Avatar src={user?.profile.avatar_url} icon={<UserOutlined />} />
          <span style={{ fontWeight: 'bold' }}>{user?.username}</span>
        </Space>
      ),
    },
    { type: 'divider' },
    {
      key: 'profile',
      icon: <UserOutlined />,
      label: 'My Profile',
    },
    {
      key: 'settings',
      icon: <SettingOutlined />,
      label: 'Settings',
    },
    {
      key: 'logout',
      icon: <LogoutOutlined />,
      label: 'Log out',
      onClick: handleLogout,
    },
  ];

  return (
    <Header style={{ 
      display: 'flex', 
      alignItems: 'center', 
      justifyContent: 'space-between', 
      backgroundColor: '#fff', 
      padding: '0 24px',
      borderBottom: '1px solid #f0f0f0',
      position: 'sticky',
      top: 0,
      zIndex: 1,
      width: '100%'
    }}>
      <div style={{ display: 'flex', alignItems: 'center' }}>
        <div className="logo" style={{ color: '#fc4c02', fontSize: '24px', fontWeight: 'bold', marginRight: '40px' }}>
          FITFEED
        </div>
        {isLoggedIn && (
          <Menu mode="horizontal" defaultSelectedKeys={['dashboard']} style={{ borderBottom: 'none', minWidth: '400px' }}>
            <Menu.Item key="dashboard">Dashboard</Menu.Item>
            <Menu.Item key="activities">Activities</Menu.Item>
            <Menu.Item key="explore">Explore</Menu.Item>
          </Menu>
        )}
      </div>
      
      <div>
        {isLoggedIn ? (
          <Dropdown menu={{ items: menuItems }} placement="bottomRight" arrow>
            <Button type="text" style={{ padding: 0, height: 'auto' }}>
              <Avatar src={user?.profile.avatar_url} icon={<UserOutlined />} />
            </Button>
          </Dropdown>
        ) : (
          <Button type="primary" style={{ backgroundColor: '#fc4c02', borderColor: '#fc4c02' }} onClick={() => setIsModalVisible(true)}>
            Log In
          </Button>
        )}
      </div>

      <Modal
        title="Log In to FitFeed"
        open={isModalVisible}
        onCancel={() => setIsModalVisible(false)}
        footer={null}
        centered
      >
        <Space direction="vertical" style={{ width: '100%' }} size="middle">
          <Button block icon={<GoogleOutlined />} onClick={() => handleOAuthLogin('google')}>
            Continue with Google
          </Button>
          <Button block icon={<IdcardOutlined />} onClick={handlePasskeyLogin}>
            Sign in with Passkey
          </Button>
          <div style={{ textAlign: 'center', color: '#8c8c8c', fontSize: '12px', marginTop: '16px' }}>
            By continuing, you agree to FitFeed's Terms of Service and Privacy Policy.
          </div>
        </Space>
      </Modal>
    </Header>
  );
};

export default AppHeader;

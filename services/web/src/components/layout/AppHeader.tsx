import React, { useState } from 'react';
import { Layout, Menu, Button, Dropdown, Avatar, Space, Modal, Input, Divider, message, type MenuProps } from 'antd';
import { UserOutlined, SettingOutlined, LogoutOutlined, GoogleOutlined, IdcardOutlined } from '@ant-design/icons';
import { useAuth } from '../../context/AuthContext';
import { passkeyService } from '../../services/auth';

const { Header } = Layout;

const AppHeader: React.FC = () => {
  const { user, isLoggedIn, logout, login, config } = useAuth();
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [username, setUsername] = useState('');
  const [loading, setLoading] = useState(false);

  const handleLogout = () => {
    logout();
  };

  const handlePasskeyLogin = async () => {
    if (!username) {
      message.error("Please enter a username");
      return;
    }
    setLoading(true);
    try {
      const userJson = await passkeyService.login(username);
      login(JSON.parse(userJson));
      message.success("Logged in with Passkey!");
      setIsModalVisible(false);
    } catch (err) {
      const error = err as Error;
      message.error(`Passkey login failed: ${error.message}`);
    } finally {
      setLoading(false);
    }
  };

  const handlePasskeyRegister = async () => {
    if (!username) {
      message.error("Please enter a username");
      return;
    }
    setLoading(true);
    try {
      await passkeyService.register(username);
      message.success("Passkey registered! You can now log in.");
    } catch (err) {
      const error = err as Error;
      message.error(`Passkey registration failed: ${error.message}`);
    } finally {
      setLoading(false);
    }
  };

  const handleOAuthLogin = (provider: string) => {
    if (!config) return;
    window.location.href = `${config.auth_url}/v1/oauth/${provider}/auth`;
  };

  const menuItems: MenuProps['items'] = [
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
    { key: 'profile', icon: <UserOutlined />, label: 'My Profile' },
    { key: 'settings', icon: <SettingOutlined />, label: 'Settings' },
    { key: 'logout', icon: <LogoutOutlined />, label: 'Log out', onClick: handleLogout },
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
          <Button type="primary" onClick={() => setIsModalVisible(true)}>
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
          <Input 
            placeholder="Username" 
            prefix={<UserOutlined />} 
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
          <Button block type="primary" icon={<IdcardOutlined />} onClick={handlePasskeyLogin} loading={loading}>
            Sign in with Passkey
          </Button>
          <Button block onClick={handlePasskeyRegister} loading={loading}>
            Register Passkey
          </Button>
          
          <Divider>Or</Divider>
          
          <Button block icon={<GoogleOutlined />} onClick={() => handleOAuthLogin('google')}>
            Continue with Google
          </Button>
        </Space>
      </Modal>
    </Header>
  );
};

export default AppHeader;

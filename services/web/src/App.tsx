import React, { useState } from 'react';
import '@ant-design/v5-patch-for-react-19';
import { Layout, Menu, Button, Modal, Dropdown, Avatar, Space, Typography, Card } from 'antd';
import { UserOutlined, SettingOutlined, LogoutOutlined, GoogleOutlined, FacebookOutlined, AppleOutlined } from '@ant-design/icons';

const { Header, Content, Footer } = Layout;
const { Title } = Typography;

// Define a type for the user object to enforce consistency
interface User {
  name: string;
  avatar?: string;
}

// Mock user data for the logged-in state
const mockUser: User = {
  name: 'Jane Doe',
  avatar: 'https://cdn.ant.design/images/ant-design-logo.svg', // Example avatar URL
};

const App: React.FC = () => {
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
  const [isModalVisible, setIsModalVisible] = useState<boolean>(false);
  const [user, setUser] = useState<User | null>(null);

  // --- Modal and Auth Handlers ---

  const handleLogin = (platform: string) => {
    console.log(`Mocking login with ${platform}...`);
    setIsLoggedIn(true);
    setUser(mockUser);
    setIsModalVisible(false); // Close modal on successful "login"
  };

  const handleLogout = () => {
    setIsLoggedIn(false);
    setUser(null);
  };

  const showModal = () => {
    setIsModalVisible(true);
  };

  const handleCancel = () => {
    setIsModalVisible(false);
  };

  // --- Dropdown Menu ---

  const profileMenu = (
    <Menu>
      <Menu.Item key="user-info" disabled style={{ padding: '8px 16px' }}>
        <Space>
          <Avatar src={user?.avatar} icon={<UserOutlined />} />
          <span style={{ fontWeight: 'bold' }}>{user?.name}</span>
        </Space>
      </Menu.Item>
      <Menu.Divider />
      <Menu.Item key="profile" icon={<UserOutlined />}>
        Profile
      </Menu.Item>
      <Menu.Item key="account" icon={<SettingOutlined />}>
        Account
      </Menu.Item>
      <Menu.Item key="logout" icon={<LogoutOutlined />} onClick={handleLogout}>
        Log out
      </Menu.Item>
    </Menu>
  );

  return (
    <Layout style={{ display: 'flex-center', minHeight: '100vh' }}>
      {/* Header */}
      <Header style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', backgroundColor: '#fff', padding: '0 24px' }}>
        <div className="logo" style={{ color: '#000', fontSize: '24px', fontWeight: 'bold' }}>
          FitFeed
        </div>
        <div>
          {isLoggedIn ? (
            <Dropdown overlay={profileMenu} placement="bottomRight" arrow>
              <Button type="text" style={{ padding: 0 }}>
                <Avatar src={user?.avatar} icon={<UserOutlined />} />
              </Button>
            </Dropdown>
          ) : (
            <Button type="primary" onClick={showModal}>
              Log In
            </Button>
          )}
        </div>
      </Header>

      {/* Main Content */}
      <Content style={{ padding: '0 50px' }}>
        <div style={{ padding: 24, minHeight: 280, marginTop: '24px', backgroundColor: '#f0f2f5', borderRadius: '8px' }}>
          <Title level={2}>Welcome to the Homepage</Title>
          <p>This is the main content area of the application.</p>
          <Card style={{ marginTop: '20px' }}>
            <p>Here you can add various components and content to your page. This is a simple mock-up to demonstrate a basic layout and state management.</p>
          </Card>
        </div>
      </Content>

      {/* Footer */}
      <Footer style={{ textAlign: 'center', backgroundColor: '#fff' }}>
        FitFeed ©2025 Created with Ant Design
      </Footer>

      {/* Login Modal */}
      <Modal
        title="Log In"
        open={isModalVisible}
        onCancel={handleCancel}
        footer={null}
        destroyOnHidden
      >
        <Space direction="vertical" style={{ width: '100%' }}>
          <Button style={{ width: '100%' }} icon={<GoogleOutlined />}  onClick={() => handleLogin('Google')}>
            Log in with Google
          </Button>
          <Button style={{ width: '100%' }} icon={<FacebookOutlined />} onClick={() => handleLogin('Facebook')}>
            Log in with Facebook
          </Button>
          <Button style={{ width: '100%' }} icon={<AppleOutlined />} onClick={() => handleLogin('Apple')}>
            Log in with Apple
          </Button>
        </Space>
      </Modal>
    </Layout>
  );
};

export default App;

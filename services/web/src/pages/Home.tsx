import React from 'react';
import { Typography, Row, Col, Card, Avatar, Button, Empty } from 'antd';
import { useAuth } from '../context/AuthContext';
import { UserOutlined, PlusOutlined } from '@ant-design/icons';

const { Title, Text } = Typography;

const Home: React.FC = () => {
  const { isLoggedIn, user } = useAuth();

  if (!isLoggedIn) {
    return (
      <div style={{ textAlign: 'center', padding: '60px 0' }}>
        <Title level={1}>The best fitfeed for your fitness journey.</Title>
        <Title level={3} type="secondary">
          Track your activities, share with friends, and stay motivated.
          Self-hosted and privacy-first.
        </Title>
        <div style={{ marginTop: '40px' }}>
          <Button type="primary" size="large" style={{ backgroundColor: '#fc4c02', borderColor: '#fc4c02', height: '50px', padding: '0 40px', fontSize: '18px' }}>
            Get Started
          </Button>
        </div>
      </div>
    );
  }

  return (
    <Row gutter={24}>
      <Col span={6}>
        <Card style={{ textAlign: 'center' }}>
          <Avatar size={64} src={user?.profile.avatar_url} icon={<UserOutlined />} />
          <Title level={4} style={{ marginTop: '16px' }}>
            {user?.profile.first_name} {user?.profile.last_name}
          </Title>
          <Text type="secondary">@{user?.username}</Text>
          <div style={{ borderTop: '1px solid #f0f0f0', marginTop: '16px', paddingTop: '16px', display: 'flex', justifyContent: 'space-around' }}>
            <div>
              <Text strong>0</Text><br />
              <Text type="secondary" style={{ fontSize: '12px' }}>FOLLOWING</Text>
            </div>
            <div>
              <Text strong>0</Text><br />
              <Text type="secondary" style={{ fontSize: '12px' }}>FOLLOWERS</Text>
            </div>
            <div>
              <Text strong>0</Text><br />
              <Text type="secondary" style={{ fontSize: '12px' }}>ACTIVITIES</Text>
            </div>
          </div>
        </Card>
      </Col>
      <Col span={12}>
        <Card title="Recent Activity" extra={<Button icon={<PlusOutlined />} type="link">Add Activity</Button>}>
          <Empty description="No activities yet. Go for a run!" />
        </Card>
      </Col>
      <Col span={6}>
        <Card title="Weekly Stats">
           <Empty image={Empty.PRESENTED_IMAGE_SIMPLE} description="No data this week" />
        </Card>
      </Col>
    </Row>
  );
};

export default Home;

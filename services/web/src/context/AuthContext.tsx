import React, { createContext, useContext, useState, type ReactNode, useEffect } from 'react';
import type { User } from '../types';

export interface AppConfig {
  auth_url: string;
  api_url: string;
}

interface AuthContextType {
  user: User | null;
  isLoggedIn: boolean;
  config: AppConfig | null;
  login: (user: User) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [user, setUser] = useState<User | null>(() => {
    const savedUser = localStorage.getItem('user');
    return savedUser ? JSON.parse(savedUser) : null;
  });
  const [config, setConfig] = useState<AppConfig | null>(null);

  useEffect(() => {
    // Fetch config from API using Vite env var for bootstrap
    const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8082';
    fetch(`${apiUrl}/v1/config`)
      .then(res => res.json())
      .then(data => setConfig(data))
      .catch(err => console.error("Failed to fetch config", err));
  }, []);

  const login = (userData: User) => {
    setUser(userData);
    localStorage.setItem('user', JSON.stringify(userData));
  };

  const logout = () => {
    setUser(null);
    localStorage.removeItem('user');
    // TODO: Call API logout using config.auth_url
  };

  return (
    <AuthContext.Provider value={{ user, isLoggedIn: !!user, config, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

// eslint-disable-next-line react-refresh/only-export-components
export const useAuth = () => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

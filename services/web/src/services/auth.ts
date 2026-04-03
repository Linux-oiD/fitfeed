import { create, get } from '@github/webauthn-json';

const getBaseUrl = () => {
  // We can get this from config in AuthContext, but for simple service we can use it here
  return 'http://localhost:8081/v1/passkey';
};

export const passkeyService = {
  async register(username: string) {
    const baseUrl = getBaseUrl();
    
    // 1. Begin Registration
    const response = await fetch(`${baseUrl}/register/begin?username=${username}`);
    if (!response.ok) throw new Error('Failed to begin registration');
    
    const options = await response.json();
    
    // 2. Create Credential
    const credential = await create(options);
    
    // 3. Finish Registration
    const finishResponse = await fetch(`${baseUrl}/register/finish?username=${username}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(credential),
    });
    
    if (!finishResponse.ok) throw new Error('Failed to finish registration');
    return await finishResponse.text();
  },

  async login(username: string) {
    const baseUrl = getBaseUrl();
    
    // 1. Begin Login
    const response = await fetch(`${baseUrl}/login/begin?username=${username}`);
    if (!response.ok) throw new Error('Failed to begin login');
    
    const options = await response.json();
    
    // 2. Get Credential
    const credential = await get(options);
    
    // 3. Finish Login
    const finishResponse = await fetch(`${baseUrl}/login/finish?username=${username}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(credential),
    });
    
    if (!finishResponse.ok) throw new Error('Failed to finish login');
    
    // The server should have set the JWT cookie at this point
    return await finishResponse.text();
  }
};

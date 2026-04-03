export interface User {
  id: string;
  username: string;
  profile: Profile;
  session?: Session;
}

export interface Profile {
  first_name: string;
  last_name: string;
  avatar_url: string;
  email: string;
}

export interface Session {
  access_token: string;
  refresh_token: string;
}

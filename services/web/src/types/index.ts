export interface User {
  id: string;
  username: string;
  profile: Profile;
}

export interface Profile {
  first_name: string;
  last_name: string;
  avatar_url: string;
  email: string;
}

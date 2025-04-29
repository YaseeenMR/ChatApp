import api, { setAuthToken } from './api';

interface User {
  name: string;
  email: string;
  password: string;
}

interface LoginData {
  email: string;
  password: string;
}

export const register = async (userData: User) => {
  try {
    const response = await api.post('/register', userData);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const login = async (loginData: LoginData) => {
  try {
    const response = await api.post('/login', loginData);
    const { token } = response.data;
    setAuthToken(token);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const getProfile = async () => {
  try {
    const response = await api.get('/profile');
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const updateProfile = async (updateData: { name?: string; password?: string }) => {
  try {
    const response = await api.patch('/profile', updateData);
    return response.data;
  } catch (error) {
    throw error;
  }
};
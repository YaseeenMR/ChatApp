// User types
export interface User {
    id?: number;
    name: string;
    email: string;
    password?: string; // Optional because we might not want to expose this in all cases
    createdAt?: string;
    updatedAt?: string;
  }
  
  // Auth types
  export interface AuthContextType {
    user: User | null;
    isLoading: boolean;
    error: string | null;
    login: (email: string, password: string) => Promise<void>;
    register: (name: string, email: string, password: string) => Promise<void>;
    logout: () => void;
  }
  
  export interface LoginData {
    email: string;
    password: string;
  }
  
  export interface RegisterData {
    name: string;
    email: string;
    password: string;
  }
  
  // Message types
  export interface Message {
    id: string;
    text: string;
    sender: 'me' | 'other';
    timestamp: Date;
  }
  
  // Navigation types
  export type AuthStackParamList = {
    Login: undefined;
    Register: undefined;
  };
  
  export type MainStackParamList = {
    Chat: undefined;
    Profile: undefined;
  };
  
  export type RootStackParamList = {
    Auth: undefined;
    Main: undefined;
  } & AuthStackParamList &
    MainStackParamList;
  
  // Component props
  export interface ButtonProps {
    title: string;
    onPress: () => void;
    style?: object;
  }
  
  export interface InputProps {
    placeholder: string;
    value: string;
    onChangeText: (text: string) => void;
    secureTextEntry?: boolean;
    style?: object;
  }
  
  export interface MessageBubbleProps {
    text: string;
    isMe: boolean;
    timestamp: Date;
  }
  
  // API response types
  export interface ApiResponse<T> {
    data?: T;
    error?: string;
    status: number;
  }
  
  export interface LoginResponse {
    token: string;
    user: User;
  }
  
  export interface RegisterResponse {
    userId: number;
    message: string;
  }
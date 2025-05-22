// Interface para o modelo de usuário
export interface User {
  id: number;
  email: string;
  name: string;
  created_at: string;
  updated_at: string; 
  deleted_at: string | null; 
}

// Interfaces para requests
export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  email: string;
  password: string;
  name: string;
}

// Interface para response de autenticação
export interface AuthResponse {
  token: string;
  user: {
    id: number;
    email: string;
    name: string;
  };
}

// Interface para erros da API
export interface ApiError {
  error: string;
} 
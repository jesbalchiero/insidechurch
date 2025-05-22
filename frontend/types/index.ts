// Interface para o modelo de usuário
export interface User {
  id: string;
  name: string;
  email: string;
  createdAt: string;
  updatedAt: string;
}

// Interfaces para requests
export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  name: string;
  email: string;
  password: string;
}

// Interface para response de autenticação
export interface AuthResponse {
  token: string;
  user: User;
}

// Interface para erros da API
export interface ApiError {
  error: string;
} 
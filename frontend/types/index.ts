// Interface para o modelo de usuário
export interface User {
  id: string;
  email: string;
}

// Interfaces para requests
export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
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

export interface LoginResponse {
  token: string;
  refreshToken: string;
}

export interface RefreshRequest {
  refresh_token: string;
}

export interface ValidateRequest {
  token: string;
}

export interface ValidateResponse {
  valid: boolean;
  user_id: string;
} 
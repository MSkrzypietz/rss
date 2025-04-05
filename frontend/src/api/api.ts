import axios from 'axios';

const baseUrl = import.meta.env.PROD ? import.meta.env.VITE_API_URL : 'http://localhost:8080';

const api = axios.create({
  headers: { 'Content-Type': 'application/json' },
  baseURL: `${baseUrl}/v1/`,
  timeout: 20000,
  withCredentials: true,
});

export default api;

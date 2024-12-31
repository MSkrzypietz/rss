import axios from 'axios';

const baseUrl = import.meta.env.PROD ? '' : 'http://localhost:8080';

const api = axios.create({
  headers: { 'Content-Type': 'application/json' },
  baseURL: `${baseUrl}/api/v1/`,
  timeout: 3000,
  withCredentials: true,
});

export default api;

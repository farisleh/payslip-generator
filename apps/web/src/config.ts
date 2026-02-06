// API configuration
const API_URL = import.meta.env.VITE_API_URL || '';

export const getApiUrl = (path: string) => {
  // In development, use relative URLs (proxy handles it)
  // In production, use the full Railway API URL
  if (API_URL) {
    return `${API_URL}${path}`;
  }
  return path;
};

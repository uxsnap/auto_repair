import axios from 'axios';

export default axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  timeout: 1000,
  transformResponse: [function (data) {
    return JSON.parse(data).data;
  }],
});

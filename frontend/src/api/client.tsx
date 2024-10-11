import axios from 'axios';

export default axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  timeout: 1000,
  transformResponse: [
    function (data) {
      const parsedData = JSON.parse(data);

      if (parsedData.error) {
        return parsedData;
      }

      return JSON.parse(data).data;
    },
  ],
});

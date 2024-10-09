import client from '../client';

export const getApps = () => {
  return client.get('/applications');
};

getApps.queryKey = 'getApps';

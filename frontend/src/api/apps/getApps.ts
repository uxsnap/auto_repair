import client from '../client';

export const getApps = () => {
  return client.get<Application[]>('/applications');
};

getApps.queryKey = 'getApps';

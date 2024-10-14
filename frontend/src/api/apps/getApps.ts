import { ApplicationFilters, ApplicationWithData, FilterValues } from '@/types';
import client from '../client';

export const getApps = (filters?: FilterValues<ApplicationFilters>) => {
  return client.get<ApplicationWithData[]>('/applications', {
    params: filters,
  });
};

getApps.queryKey = 'getApps';

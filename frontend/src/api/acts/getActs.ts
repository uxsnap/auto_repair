import { ActFilters, ActWithData, FilterValues } from '@/types';
import client from '../client';

export const getActs = (filters?: FilterValues<ActFilters>) => {
  return client.get<ActWithData[]>('/acts', {
    params: filters,
  });
};

getActs.queryKey = 'getActs';

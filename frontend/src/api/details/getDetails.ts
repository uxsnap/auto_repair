import { Detail, FilterValues } from '@/types';
import client from '../client';

export const getDetails = (filters: FilterValues<Detail>) => {
  return client.get<Detail[]>('/details', {
    params: filters
  });
};

getDetails.queryKey = 'getDetails';

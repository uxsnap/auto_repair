import { Detail, DetailsFilters, FilterValues } from '@/types';
import client from '../client';

export const getDetails = (filters?: FilterValues<DetailsFilters>) => {
  return client.get<Detail[]>('/details', {
    params: filters,
  });
};

getDetails.queryKey = 'getDetails';

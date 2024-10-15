import { Contract, FilterValues } from '@/types';
import client from '../client';

export const getContracts = (filters?: FilterValues<any>) => {
  return client.get<Contract[]>('/contracts', {
    params: filters,
  });
};

getContracts.queryKey = 'getContracts';

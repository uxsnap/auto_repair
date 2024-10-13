import { ClientWithData, FilterValues, Storage, StorageWithData } from '@/types';
import client from '../client';

export const getClients = (filters: FilterValues<any>) => {
  return client.get<ClientWithData[]>('/clients', {
    params: filters,
  });
};

getClients.queryKey = 'getClients';

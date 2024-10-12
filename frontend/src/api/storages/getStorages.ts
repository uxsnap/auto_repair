import { FilterValues, Storage, StorageFilters, StorageWithData } from '@/types';
import client from '../client';

export const getStorages = (filters: FilterValues<StorageFilters>) => {
  return client.get<StorageWithData[]>('/storages', {
    params: filters,
  });
};

getStorages.queryKey = 'getStorages';

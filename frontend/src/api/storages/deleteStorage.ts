import { IdBody } from '@/types';
import client from '../client';

export const deleteStorage = (body: IdBody) => {
  return client.delete('/storages', {
    data: body,
  });
};

deleteStorage.queryKey = 'deleteStorage';

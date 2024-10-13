import { IdBody } from '@/types';
import client from '../client';

export const deleteClient = (body: IdBody) => {
  return client.delete('/clients', {
    data: body,
  });
};

deleteClient.queryKey = 'deleteClient';

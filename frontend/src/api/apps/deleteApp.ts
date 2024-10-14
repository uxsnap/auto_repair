import { IdBody } from '@/types';
import client from '../client';

export const deleteApp = (body: IdBody) => {
  return client.delete('/applications', {
    data: body,
  });
};

deleteApp.queryKey = 'deleteApp';

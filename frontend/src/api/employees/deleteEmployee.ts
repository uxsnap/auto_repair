import { IdBody } from '@/types';
import client from '../client';

export const deleteEmployee = (body: IdBody) => {
  return client.delete('/employees', {
    data: body,
  });
};

deleteEmployee.queryKey = 'deleteEmployee';

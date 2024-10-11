import { IdBody } from '@/types';
import client from '../client';

export const deleteDetail = (body: IdBody) => {
  return client.delete('/details', {
    data: body,
  });
};

deleteDetail.queryKey = 'deleteDetail';

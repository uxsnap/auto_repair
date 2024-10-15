import { IdBody } from '@/types';
import client from '../client';

export const deleteContract = (body: IdBody) => {
  return client.delete('/contracts', {
    data: body,
  });
};

deleteContract.queryKey = 'deleteContract';

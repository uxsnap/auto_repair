import { IdBody } from '@/types';
import client from '../client';

export const deleteAct = (body: IdBody) => {
  return client.delete('/acts', {
    data: body,
  });
};

deleteAct.queryKey = 'deleteAct';

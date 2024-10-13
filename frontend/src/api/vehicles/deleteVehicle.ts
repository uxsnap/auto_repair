import { IdBody } from '@/types';
import client from '../client';

export const deleteVehicle = (body: IdBody) => {
  return client.delete('/vehicles', {
    data: body,
  });
};

deleteVehicle.queryKey = 'deleteVehicle';

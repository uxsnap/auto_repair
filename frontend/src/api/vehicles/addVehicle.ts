import { IdBody } from '@/types';
import client from '../client';

type AddVehicleBody = {
  ClientId: string;
  VehicleNumber: string;
  Brand: string;
  Model: string;
};

export const addVehicle = (body: AddVehicleBody) => {
  return client.post<IdBody[]>('/vehicles', body);
};

addVehicle.queryKey = 'addVehicle';

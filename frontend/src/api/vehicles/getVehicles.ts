import { FilterValues, VehiclesFilters, VehicleWithData } from '@/types';
import client from '../client';

export const getVehicles = (filters: FilterValues<VehiclesFilters>) => {
  return client.get<VehicleWithData[]>('/vehicles', {
    params: filters,
  });
};

getVehicles.queryKey = 'getVehicles';

import { Service } from '@/types';
import client from '../client';

export const getServices = () => {
  return client.get<Service[]>('/services');
};

getServices.queryKey = 'getServices';

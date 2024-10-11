import { Employee } from '@/types';
import client from '../client';

export const getEmployees = () => {
  return client.get<Employee[]>('/employees');
};

getEmployees.queryKey = 'getEmployees';

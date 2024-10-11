import { IdBody } from '@/types';
import client from '../client';

type AddEmployeeBody = {
  Name: string;
  Position: string;
  EmployeeNum: string;
};

export const addEmployee = (body: AddEmployeeBody) => {
  return client.post<IdBody[]>('/employees', body);
};

addEmployee.queryKey = 'addEmployee';

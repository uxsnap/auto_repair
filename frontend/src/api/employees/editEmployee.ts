import { IdBody } from '@/types';
import client from '../client';

type EditEmployeeBody = {
  Id: string;
  Name: string;
  Position: string;
  EmployeeNum: string;
};

export const editEmployee = (body: EditEmployeeBody) => {
  return client.patch<IdBody>(`/employees/${body.Id}`, body);
};

editEmployee.queryKey = 'editEmployee';

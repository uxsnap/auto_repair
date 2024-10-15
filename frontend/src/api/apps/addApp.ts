import { IdBody } from '@/types';
import client from '../client';

type AddAppBody = {
  EmployeeId: string;
  ClientId: string;
  ContractId: string;
  Name: string;
  Status: string;
};

export const addApp = (body: AddAppBody) => {
  console.log(body)

  return client.post<IdBody[]>('/applications', body);
};

addApp.queryKey = 'addApp';

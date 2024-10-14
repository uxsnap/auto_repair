import { IdBody } from '@/types';
import client from '../client';

type EditAppBody = {
  Id: string;
  EmployeeId: string;
  ClientId: string;
  Name: string;
  Status: string;
  ContractId: string;
};

export const editApp = (body: EditAppBody) => {
  return client.patch<IdBody>(`/applications/${body.Id}`, body);
};

editApp.queryKey = 'editApp';

import { IdBody } from '@/types';
import client from '../client';

type AddStorageBody = {
  EmployeeId: string;
  DetailId: string;
  StorageNum: string;
  DetailCount: number;
};

export const addStorage = (body: AddStorageBody) => {
  return client.post<IdBody[]>('/storages', body);
};

addStorage.queryKey = 'addStorage';

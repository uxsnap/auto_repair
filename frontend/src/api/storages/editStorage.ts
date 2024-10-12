import { IdBody } from '@/types';
import client from '../client';

type EditStorageBody = {
  Id: string;
  EmployeeId: string;
  DetailId: string;
  StorageNum: string;
  DetailCount: number;
};

export const editStorage = (body: EditStorageBody) => {
  return client.patch<IdBody>(`/storages/${body.Id}`, body);
};

editStorage.queryKey = 'editStorage';

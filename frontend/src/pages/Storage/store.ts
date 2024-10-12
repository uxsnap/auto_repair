import { create } from 'zustand';
import { FilterValues, StorageFilters } from '@/types';

type Store = {
  filters: FilterValues<StorageFilters>;
  setFilters: (payload: FilterValues<StorageFilters>) => void;
};

export const useFiltersStore = create<Store>()((set) => ({
  filters: {
    storageNum: '',
  },
  setFilters: (payload: FilterValues<StorageFilters>) => set(() => ({ filters: payload })),
}));

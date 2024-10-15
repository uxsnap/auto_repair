import { create } from 'zustand';
import { ContractsFilters, FilterValues } from '@/types';

type Store = {
  filters: FilterValues<ContractsFilters>;
  setFilters: (payload: FilterValues<ContractsFilters>) => void;
};

export const useFiltersStore = create<Store>()((set) => ({
  filters: {
    name: '',
  },
  setFilters: (payload: FilterValues<ContractsFilters>) => set(() => ({ filters: payload })),
}));

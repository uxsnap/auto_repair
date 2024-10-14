import { create } from 'zustand';
import { ApplicationFilters, FilterValues } from '@/types';

type Store = {
  filters: FilterValues<ApplicationFilters>;
  setFilters: (payload: FilterValues<ApplicationFilters>) => void;
};

export const useFiltersStore = create<Store>()((set) => ({
  filters: {
    name: '',
  },
  setFilters: (payload: FilterValues<ApplicationFilters>) => set(() => ({ filters: payload })),
}));

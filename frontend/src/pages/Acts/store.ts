import { create } from 'zustand';
import { ActFilters, FilterValues } from '@/types';

type Store = {
  filters: FilterValues<ActFilters>;
  setFilters: (payload: FilterValues<ActFilters>) => void;
};

export const useFiltersStore = create<Store>()((set) => ({
  filters: {
    name: '',
  },
  setFilters: (payload: FilterValues<ActFilters>) => set(() => ({ filters: payload })),
}));

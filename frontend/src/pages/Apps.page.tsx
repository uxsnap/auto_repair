import { useQuery } from '@tanstack/react-query';
import { getApps } from '@/api/apps/getApps';

export function AppsPage() {
  const { data } = useQuery({ queryKey: [getApps.queryKey], queryFn: getApps });

  console.log(data);

  return <>Заявки</>;
}

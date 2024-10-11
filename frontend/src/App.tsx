import '@mantine/core/styles.css';

import { IconEngine } from '@tabler/icons-react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import {
  createBrowserRouter,
  Route,
  RouterProvider,
  Routes,
  useLocation,
  useNavigate,
} from 'react-router-dom';
import { AppShell, MantineProvider, Tabs, Title } from '@mantine/core';
import { Notifications } from '@mantine/notifications';
import { AppsPage } from './pages/Apps/Apps.page';
import { DetailsPage } from './pages/Details/Details.page';
import { theme } from './theme';

import '@mantine/notifications/styles.css';

const queryClient = new QueryClient();

const router = createBrowserRouter([{ path: '*', element: <Root /> }]);

function Root() {
  const navigate = useNavigate();
  const location = useLocation();

  return (
    <MantineProvider theme={theme}>
      <Notifications />

      <QueryClientProvider client={queryClient}>
        <AppShell header={{ height: 60 }} padding="md">
          <AppShell.Header
            display="flex"
            style={{ alignItems: 'center', gap: 12 }}
            px={20}
            c="blue"
          >
            <IconEngine size={40} />
            <Title order={3}>Auto Repair</Title>
          </AppShell.Header>

          <AppShell.Main>
            <Tabs
              defaultValue={location.pathname}
              onChange={(value) => navigate(`${value === '/' ? '' : value}`)}
              value={location.pathname}
              variant="outline"
            >
              <Tabs.List>
                <Tabs.Tab value="/">Заявки</Tabs.Tab>
                <Tabs.Tab value="/clients">Клиенты</Tabs.Tab>
                <Tabs.Tab value="/details">Детали</Tabs.Tab>
              </Tabs.List>

              <Routes>
                <Route path="/clients" element={<AppsPage />} />
                <Route path="/details" element={<DetailsPage />} />
                <Route path="/" element={<AppsPage />} />
              </Routes>
            </Tabs>
          </AppShell.Main>
        </AppShell>
      </QueryClientProvider>
    </MantineProvider>
  );
}

export default function App() {
  return <RouterProvider router={router} />;
}

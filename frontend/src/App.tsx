import '@mantine/core/styles.css';

import { IconEngine } from '@tabler/icons-react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import {
  createBrowserRouter,
  Route,
  RouterProvider,
  Routes,
  useNavigate,
  useParams,
} from 'react-router-dom';
import { AppShell, Button, Group, MantineProvider, Tabs, Title } from '@mantine/core';
import { AppsPage } from './pages/Apps/Apps.page';
import { theme } from './theme';

const queryClient = new QueryClient();

const router = createBrowserRouter([{ path: '*', element: <Root /> }]);

function Root() {
  const navigate = useNavigate();
  const { tabValue } = useParams();

  return (
    <MantineProvider theme={theme}>
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
              defaultValue="/"
              onChange={(value) => navigate(`/${value === '/' ? '' : value}`)}
              value={tabValue}
              variant='outline'
            >
              <Tabs.List>
                <Tabs.Tab value="/">Заявки</Tabs.Tab>
                <Tabs.Tab value="clients">Клиенты</Tabs.Tab>
              </Tabs.List>

              <Routes>
                <Route path="/" element={<AppsPage />} />
                <Route path="/clients" element={<AppsPage />} />
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

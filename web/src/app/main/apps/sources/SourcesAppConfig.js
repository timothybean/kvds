import { lazy } from 'react';
import { Navigate } from 'react-router-dom';
import SourcesApp from './SourcesApp';
import SourceView from './source/SourceView';
import Sources from './sources/Sources'

const Course = lazy(() => import('./source/Source'));
const Courses = lazy(() => import('./sources/Sources'));

const SourcesAppConfig = {
  settings: {
    layout: {},
  },
  routes: [
    {
      path: 'apps/sources',
      element: <SourcesApp />,
      children: [
        {
          path: '',
          element: <Sources />,  //<Navigate to="/apps/sources/sources" />,
        },
        {
          path: ':id',
          element: <SourceView />,
        },
        {
          path: 'sources',
          element: <Courses />,
        },
      ],
    },
  ],
};

export default SourcesAppConfig;

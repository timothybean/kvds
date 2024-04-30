import { Outlet } from 'react-router-dom';
import FusePageSimple from '@fuse/core/FusePageSimple';
import { useEffect, useRef, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useParams } from 'react-router-dom';
import withReducer from 'app/store/withReducer';
import { styled } from '@mui/material/styles';
import { getCategories } from './store/categoriesSlice';
import useThemeMediaQuery from '@fuse/hooks/useThemeMediaQuery';
import reducer from './store';
import SourceSidebarContent from './sources/SourceSidebarContent';
import Sources from './sources/Sources';
import { getWidgets, selectWidgets } from './store/widgetsSlice';

const Root = styled(FusePageSimple)(({ theme }) => ({
  '& .FusePageSimple-header': {
    backgroundColor: theme.palette.background.paper,
  },
}));

function SourcesApp() {
  const widgets = useSelector(selectWidgets);
  const dispatch = useDispatch();
  const pageLayout = useRef(null);
  
  useEffect(() => {
    dispatch(getCategories());
  }, [dispatch]);

  useEffect(() => {
    dispatch(getWidgets());
  }, [dispatch]);

/*   if (_.isEmpty(widgets)) {
    return null;
  } */

  /* return <Outlet />; */

  return (
    <Root
      content={<Sources />}
      ref={pageLayout}
    />
  );
}

export default SourcesApp;

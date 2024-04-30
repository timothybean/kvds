import { combineReducers } from '@reduxjs/toolkit';
import course from './courseSlice';
import courses from './coursesSlice';
import categories from './categoriesSlice';
import widgets from './widgetsSlice'

const reducer = combineReducers({
  widgets,
  categories,
  courses,
  course,
});

export default reducer;

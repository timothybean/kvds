import { darken, lighten } from '@mui/material/styles';
import Chip from '@mui/material/Chip';
import { useSelector } from 'react-redux';
import _ from '@lodash';
import { selectCategories } from './store/categoriesSlice';

function CourseCategory({ id }) {
  const categories = useSelector(selectCategories);
console.log(categories)
  const category = _.find(categories, { id });

  return (
    <Chip
      className="font-semibold text-12"
      label={category?.title}
      sx={{
        color: (theme) =>
          theme.palette.mode === 'light'
            ? darken(category?.color, 0.4)
            : lighten(category?.color, 0.8),
        backgroundColor: (theme) =>
          theme.palette.mode === 'light'
            ? lighten(category?.color, 0.8)
            : darken(category?.color, 0.7),
      }}
      size="small"
    />
  );
}

export default CourseCategory;

import {
  faCode,
  faHome,
} from '@fortawesome/free-solid-svg-icons';

const tabList = [
  {
    path: '/',
    pageName: 'ホーム',
    icon: faHome,
  },
  {
    path: '/clubs/1',
    pageName: 'プログラミング研究部',
    icon: faCode,
  },
] as const;

export default tabList;

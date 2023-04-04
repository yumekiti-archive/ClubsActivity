import { FC } from 'react';
import {
  faSearch,
  faUserCheck,
} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import HomeItem from 'components/molecules/HomeItem';


const ShortcutCard: FC = () => {
  return (
    <div className='w-full shadow p-5 rounded-lg bg-white mb-2'>
      <div className='flex justify-center items-center grid grid-cols-2 mb-4'>
        <p className='text-center text-gray-700'>04/02</p>
        <select id="example8" className="text-gray-500 block w-full border-0 border-b-2 border-gray-200 focus:border-primary-500 focus:ring-0 disabled:cursor-not-allowed">
          <option value="">Underline</option>
          <option value="">Option01</option>
          <option value="">Option02</option>
          <option value="">Option03</option>
        </select>
      </div>
      <div className='grid grid-cols-2 gap-4 md:gap-8 xl:grid-cols-2 px-2'>
        <HomeItem
          title='詳細'
          icon={faSearch}
          path='/class'
        >
          選択している部活の詳細を確認できます。
        </HomeItem>
        <HomeItem
          title='出席'
          icon={faUserCheck}
          path='/class'
        >
          選択している部活に出席できます。
        </HomeItem>
      </div>
    </div>
  );
};

export default ShortcutCard;
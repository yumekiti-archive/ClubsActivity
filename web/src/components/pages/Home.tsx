import { FC } from 'react';
import Layout from 'components/template/Layout';
import Title from 'components/template/Title';
import ProfileCard from 'components/organisms/ProfileCard';
import ShortcutCard from 'components/organisms/ShortcutCard';

const Home: FC = () => {
  return (
    <Layout>
      <Title pageTitle='ホーム画面' />
      <div className='grid xl:grid-cols-2 gap-4 mx-auto'>
        <ProfileCard />
        <ShortcutCard />
      </div>
    </Layout>
  );
};

export default Home;

import React, { VFC } from 'react';
import Layout from 'components/template/Layout';
import Title from 'components/template/Title';

const Home: VFC = () => {
  return (
    <Layout>
      <Title pageTitle='ホーム画面' />
      <div className='grid grid-cols-2 gap-4 md:gap-8 xl:grid-cols-3 px-2 pb-24'>
        hoge
      </div>
    </Layout>
  );
};

export default Home;

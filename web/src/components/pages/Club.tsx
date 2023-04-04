import { FC } from 'react';
import Layout from 'components/template/Layout';
import Title from 'components/template/Title';

const Club: FC = () => {
  return (
    <Layout>
      <Title pageTitle='ホーム画面' />
      <div>
        <h1>部活詳細</h1>
      </div>
    </Layout>
  );
};

export default Club;
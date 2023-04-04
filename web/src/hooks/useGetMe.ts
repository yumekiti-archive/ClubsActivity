import User from 'types/user';
import { useQuery } from 'react-query';
import { fetchInstance } from 'libs/fetchInstance';

const getMe = async () => {
  const { data } = await fetchInstance().get('/me');
  return data;
};

const useGetMe = () => {
  const queryFn = () => getMe();
  return useQuery({
    queryKey: 'me',
    queryFn,
    cacheTime: 60 * 60 * 1000,
    staleTime: 60 * 60 * 1000,
  });
};

export default useGetMe;
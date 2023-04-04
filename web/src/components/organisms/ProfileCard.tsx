import { VFC } from 'react'

const ProfileCard: VFC = () => {
  return (
    <div className='w-full shadow p-5 rounded-lg bg-white mb-2 flex items-center justify-center'>
      <div className='w-full h-full grid grid-cols-3'>
        <div className='col-span-1 flex items-center justify-center flex-col'>
          <img
            className='h-24 w-24 rounded-full'
            src='https://source.unsplash.com/random'
            alt='avatar'
          />
          <div>
            <h2 className='text-2xl font-bold text-gray-700'>
              John Doe
            </h2>
            <p className='text-gray-500'>
              IE2A
            </p>
          </div>
        </div>
        <div className='col-span-2'>
          <div className='flex justify-end mb-8'>
          <button className="rounded-full border border-blue-500 bg-blue-500 px-5 py-2.5 text-center text-sm font-medium text-white shadow-sm transition-all hover:border-blue-700 hover:bg-blue-700 focus:ring focus:ring-blue-200 disabled:cursor-not-allowed disabled:border-blue-300 disabled:bg-blue-300">
            プロフィール
          </button>
          </div>
          <div className='flex justify-center items-center flex-col'>
            <h2 className='text-2xl font-bold text-gray-500'>
              所属
            </h2>
            <p className='text-xl text-gray-700'>
              ネットワーク研究会
            </p>
          </div>
        </div>
      </div>
    </div>
  )
}

export default ProfileCard